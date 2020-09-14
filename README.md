# evcc

[![Build Status](https://travis-ci.org/andig/evcc.svg?branch=master)](https://travis-ci.org/andig/evcc)
[![Code Quality](https://goreportcard.com/badge/github.com/andig/evcc)](https://goreportcard.com/report/github.com/andig/evcc)
[![Latest Version](https://img.shields.io/github/tag/andig/evcc.svg)](https://github.com/andig/evcc/releases)
[![Pulls from Docker Hub](https://img.shields.io/docker/pulls/andig/evcc.svg)](https://hub.docker.com/r/andig/evcc)


EVCC is an extensible EV Charge Controller with PV integration implemented in [Go](2).

## Features

- simple and clean responsive web user interface to be used on phone, tablet or pc
- multiple [chargers](#charger) including Wallbe, Phoenix (includes ESL Walli), go-eCharger, NRGkick (direct Bluetooth or via Connect device), SimpleEVSE, EVSEWifi, KEBA/BMW, openWB, Mobile Charger Connect and any other charger using scripting
- multiple [meters](#meter) including ModBus devices (Eastron SDM, MPM3PM, SBC ALE3, ORNO and many more), Discovergy (using HTTP plugin), SMA Home Manager 2.0 and SMA Energy Meter, KOSTAL Smart Energy Meter (KSEM, EMxx), any other Sunspec-compatible inverter or home battery devices (Fronius, SMA, SolarEdge, KOSTAL, STECA, E3DC, ...), Tesla PowerWall
- different [vehicles](#vehicle) to show battery status and charge time estimation like Audi (eTron), BMW (i3), Tesla, Nissan (Leaf), Renault ZE (ZOE, ...), Kia, Hyundai, VW-brands, Porsche and any other vehicle using scripting
- [plugins](#plugins) for integrating with other hardware devices and home automation: Modbus (meters and grid inverters), MQTT and shell scripts
- YAML configuration allows for flexible mixture of infrastructure from any vendor
- status notifications using [Telegram](https://telegram.org) and [PushOver](https://pushover.net)
- logging using [InfluxDB](https://www.influxdata.com) and [Grafana](https://grafana.com/grafana/)
- soft ramp-up/ramp-down of charge current ensures contactor only switched at minimum current
- REST API and MQTT
- single executable and single configuration file

![Screenshot](docs/screenshot.png)

## Index

- [Getting started](#getting-started)
- [Installation](#installation)
- [Configuration](#configuration)
  - [Site](#site)
  - [Loadpoint](#loadpoint)
  - [Charger](#charger)
  - [Meter](#meter)
  - [Vehicle](#vehicle)
- [Plugins](#plugins)
  - [Modbus](#modbus-read-only)
  - [MQTT](#mqtt-readwrite)
  - [Script](#script-readwrite)
  - [HTTP](#http-readwrite)
  - [Websocket](#websocket-read-only)
  - [Combined status](#combined-status-read-only)
- [API](#api)
- [Background](#background)

## Getting started

1. Install EVCC. For details see [installation](#installation).
2. Copy the default configuration file `evcc.dist.yaml` to `evcc.yaml` and open for editing.
3. To create a minimal setup you need a [meter](#meter) (either grid meter or at least one pv generation meter) and a supported [charger](#charger). Both need be combined to a [loadpoint](#loadpoint).
4. Configure meter(s) and charger by:
    - choosing the appropriate `type`
    - add a `name` attribute than can later be referred to
    - add configuration details depending on `type`
  See `evcc.dist.yaml` for examples.
5. Test your meter, charger and optional vehicle configuration by running

        evcc meter
		evcc charger
		evcc vehicle

6. Configure a loadpoint and refer to the meter, charger and vehicle using the defined `name` attributes.
7. Provide optional configuration for MQTT, push messaging, database logging and custom menus.

## Installation

EVCC is provided as binary executable file and docker image. Download the file for your platform and then execute like this:

    evcc -h

Use the following `systemd` unit description to configure EVCC as service (put into `/etc/systemd/system/evcc.service`):

    [Unit]
    Description=evcc
    After=syslog.target
    [Service]
    ExecStart=/usr/local/bin/evcc --log error
    Restart=always
    [Install]
    WantedBy=multi-user.target

EVCC can also be run using Docker. Here's and example with given config file and UI on port 7070:

```sh
docker run -v $(pwd)/evcc.dist.yaml:/etc/evcc.yaml -p 7070:7070 andig/evcc -h
```

If using Docker with a meter or charger that requires UDP like KEBA, make sure that the Docker container can receive UDP messages on the relevant ports (`:7090` for KEBA):

```sh
docker run -p 7070:7070 -p 7090:7090/udp andig/evcc ...
```

When using Docker with a device that requires multicast UDP like SMA, make sure that the Docker container uses the `network_mode: host` configuration:

```sh
docker run --network host andig/evcc ...
```

To build EVCC from source, [Go](2) 1.13 is required:

    make

**Note**: EVCC comes without any guarantee. You are using this software **entirely** at your own risk. It is your responsibility to verify it is working as intended.
EVCC requires a supported charger and a combination of grid, PV and charge meter.
All components **must** be installed by a certified professional.

## Configuration

The EVCC consists of five basic elements: *Site* and *Loadpoints* describe the infrastructure and combine *Charger*s, *Meter*s and *Vehicle*s.

### Site

A site describes the grid connection and is responsible for managing the available power. A minimal site configuration requires a grid meter for managing EVU demand and optionally a PV or battery meter.

```yaml
site:
- title: Zuhause # display name for UI
  meters:
    grid: sdm630 # grid meter reference
    pv: sma # pv meter reference
```

### Loadpoint

Loadpoints combine meters, charger and vehicle together and add optional configuration. A minimal loadpoint configuration requires a charger and optionally a separate charge meter. If charger has an integrated meter it will automatically be used:

```yaml
loadpoints:
- title: Garage # display name for UI
  charger: wallbe # charger reference
  vehicle: audi # vehicle reference
  meters:
    charge: sdm630 # grid meter reference
```

More options are documented in the `evcc.dist.yaml` sample configuration.

#### Charge modes

The default *charge mode* upon start of EVCC is configured on the loadpoint. Multiple charge modes are supported:

- **Off**: disable the charger, even if car gets connected.
- **Now** (**Sofortladen**): charge immediately with maximum allowed current.
- **Min + PV**: charge immediately with minimum configured current. Additionally use PV if available.
- **PV**: use PV as available. May not charge the car if PV remains dark.

In general, due to the minimum value of 5% for signalling the EV duty cycle, the charger cannot limit the current to below 6A. If the available power calculation demands a limit less than 6A, handling depends on the charge mode. In **PV** mode, the charger will be disabled until available PV power supports charging with at least 6A. In **Min + PV** mode, charging will continue at minimum current of 6A and charge current will be raised as PV power becomes available again.

### Charger

Charger is responsible for handling EV state and adjusting charge current. Available charger implementations are:

- `wallbe`: Wallbe Eco chargers (see [Preparation](#wallbe-preparation)). For older Wallbe boxes (pre 2019) with Phoenix EV-CC-AC1-M3-CBC-RCM-ETH controllers make sure to set `legacy: true` to enable correct current configuration.
- `phoenix-emcp`: chargers with Phoenix EM-CP-PP-ETH controllers like the ESL Walli (Ethernet connection).
- `phoenix-evcc`: chargers with Phoenix EV-CC-AC1-M controllers (ModBus connection)
- `simpleevse`: chargers with SimpleEVSE controllers connected via ModBus (e.g. OpenWB Wallbox, Easy Wallbox B163, ...)
- `evsewifi`: chargers with SimpleEVSE controllers using [EVSE-WiFi](https://www.evse-wifi.de/)
- `nrgkick-bluetooth`: NRGkick chargers with Bluetooth connector (Linux only, not supported on Docker)
- `nrgkick-connect`: NRGkick chargers with additional NRGkick Connect module
- `go-e`: go-eCharger chargers (both local and cloud API are supported)
- `keba`: KEBA KeContact P20/P30 and BMW chargers (see [Preparation](#keba-preparation))
- `mcc`: Mobile Charger Connect devices (Audi, Bentley, Porsche)
- `default`: default charger implementation using configurable [plugins](#plugins) for integrating any type of charger

Configuration examples are documented at [andig/evcc-config#chargers](https://github.com/andig/evcc-config#chargers)

#### Wallbe preparation

Wallbe chargers are supported out of the box. The Wallbe must be connected using Ethernet. If not configured, the default address `192.168.0.8:502` is used.

To allow controlling charge start/stop, the Wallbe physical configuration must be modified. This requires opening the Wallbe. Once opened, DIP 10 must be set to ON:

![dip10](docs/dip10.jpeg)

More information on interacting with Wallbe chargers can be found at [GoingElectric](https://www.goingelectric.de/forum/viewtopic.php?p=1212583). Use with care.

**NOTE:** The Wallbe products come in two flavors. Older models (2017 known to be "old", 2019 known to be "new") use the Phoenix EV-CC-AC1-M3-CBC-RCM controller. For such models make sure to set `legacy: true`. You can find you which one you have using [MBMD](5):

```sh
mbmd read -a 192.168.0.8:502 -d 255 -t holding -e int 300 1
```

Compare the value to what you see as *Actual Charge Current Setting* in the Wallbe web UI. If the numbers match, it's a Phoenix controller, if the reading is factor 10x the UI value then it's a Wallbe controller.

**NOTE:** Opening the wall box **must** only be done by certified professionals. The box **must** be disconnected from mains before opening.

#### KEBA preparation

KEBA chargers require UDP function to be enabled with DIP switch 1.3 = `ON`, see KEBA installation manual.

### Meter

Meters provide data about power and energy consumption or PV production. Available meter implementations are:

- `modbus`: ModBus meters as supported by [MBMD](https://github.com/volkszaehler/mbmd#supported-devices). Configuration is similar to the [ModBus plugin](#modbus-read-only) where `power` and `energy` specify the MBMD measurement value to use.
- `sma`: SMA Home Manager 2.0 and SMA Energy Meter. Power reading is configured out of the box but can be customized if necessary. To obtain specific energy readings define the desired Obis code (Import Energy: "1:1.8.0", Export Energy: "1:2.8.0").
- `tesla`: Tesla PowerWall meter. Use `usage` to choose meter (grid meter: `site`, pv: `solar`, battery: `battery`).
  *Note*: this could also be implemented using a `default` meter with the `http` plugin.
- `default`: default meter implementation where meter readings- `power` and `energy` are configured using [plugins](#plugins)

Configuration examples are documented at [andig/evcc-config#meters](https://github.com/andig/evcc-config#meters)

### Vehicle

Vehicle represents a specific EV vehicle and its battery. If vehicle is configured and assigned to the charger, charge status and remaining charge duration become available in the user interface.

Available vehicle implementations are:

- `audi`: Audi (eTron)
- `bmw`: BMW (i3)
- `nissan`: Nissan (Leaf)
- `tesla`: Tesla (any model)
- `renault`: Renault (Zoe, Kangoo ZE)
- `porsche`: Porsche (Taycan)
- `kia`: Kia (Bluelink vehicles like Soul 2019)
- `hyundai`: Hyundai (Bluelink vehicles like Kona or Ioniq)
- `default`: default vehicle implementation using configurable [plugins](#plugins) for integrating any other type of vehicle

Configuration examples are documented at [andig/evcc-config#vehicles](https://github.com/andig/evcc-config#vehicles)

## Plugins

Plugins are used to integrate various devices and external data sources with EVCC. Plugins can be used in combination with a `default` type meter, charger or vehicle.

Plugins support both *read* and *write* access. When using plugins for *write* access, the actual data is provided as variable in form of `${var[:format]}`. If `format` is omitted, data is formatted according to the default Go `%v` [format](https://golang.org/pkg/fmt/). The variable is replaced with the actual data before the plugin is executed.

### Calc (read only)

The `calc` plugin allows calculating the sum of other plugins:

```yaml
type: calc
add:
- type: ...
  ...
- type: ...
  ...
```

The `calc` plugin is useful e.g. to combine power values if import and export power are separate like with S0 meters. Use `scale` on one of the elements to implement a subtraction.

### Modbus (read only)

The `modbus` plugins is able to read data from any Modbus meter or SunSpec-compatible solar inverter. Many meters are already pre-configured (see [MBMD Supported Devices](https://github.com/volkszaehler/mbmd#supported-devices)).

The meter configuration consists of the actual physical connection and the value to be read.

#### Physical connection

If the device is physically connected using an RS485 adapter, `device` and serial configuration `baudrate`, `comset` must be specified:

```yaml
type: modbus
device: /dev/ttyUSB0
baudrate: 9600
comset: "8N1"
```

If the device is a grid inverter or a Modbus meter connected via TCP, `uri` must be specified:

```yaml
type: modbus
uri: 192.168.0.11:502
id: 1 # modbus slave id
```

If the device is a Modbus RTU device connected using an RS485/Ethernet adapter, set `rtu: true`. The serial configuration must be done directly on the adapter. Example:

```yaml
type: modbus
uri: 192.168.0.10:502
id: 3 # modbus slave id
rtu: true
```

#### Logical connection

The meter device type `meter` and the device's slave id `id` are always required:

```yaml
type: ...
uri/device/id: ...
model: sdm
value: Power
scale: -1 # floating point factor applied to result, e.g. for kW to W conversion
```

Supported meter models are the same as supported by [MBMD](https://github.com/volkszaehler/mbmd#supported-devices):

- RTU:
  - `ABB` ABB A/B-Series meters
  - `MPM`  Bernecker Engineering MPM3PM meters
  - `DZG` DZG Metering GmbH DVH4013 meters
  - `INEPRO` Inepro Metering Pro 380
  - `JANITZA`  Janitza B-Series meters
  - `SBC` Saia Burgess Controls ALD1 and ALE3 meters
  - `SDM` Eastron SDM630
  - `SDM220` Eastron SDM220
  - `SDM230` Eastron SDM230
  - `SDM72` Eastron SDM72
  - `ORNO1P` ORNO WE-514 & WE-515
  - `ORNO3P` ORNO WE-516 & WE-517
- TCP: Sunspec-compatible grid inverters (SMA, SolarEdge, KOSTAL, Fronius, Steca etc, ...)

Use `value` to define the value to read from the device. All values that are supported by [MBMD](https://github.com/volkszaehler/mbmd/blob/master/meters/measurements.go#L28) are pre-configured.

In case of SunSpec-compatible inverters, values can also be configured in the form of `model:[block:]point` according to SunSpec definition. For example, a 3-phase inverter's DC power of the 2nd string would be configurable as `103:2:W`.

#### Manual configuration

If the Modbus device is not supported by MBMD, the Modbus register can also be manually configured:

```yaml
type: ...
uri/device/id: ...
register:
  address: 40070
  type: holding # holding or input
  decode: int32 # int16|32|64, uint16|32|64, float32|64 and u|int32s + float32s
scale: -1 # floating point factor applied to result, e.g. for kW to W conversion
```

The `int32s/uint32s` decodings apply swapped word order and are useful e.g. with E3/DC devices.

### MQTT (read/write)

The `mqtt` plugin allows to read values from MQTT topics. This is particularly useful for meters, e.g. when meter data is already available on MQTT. See [MBMD](5) for an example how to get Modbus meter data into MQTT.

Sample configuration:

```yaml
type: mqtt
topic: mbmd/sdm1-1/Power
timeout: 30s # don't accept values older than timeout
scale: 0.001 # floating point factor applied to result, e.g. for Wh to kWh conversion
```

Sample write configuration:

```yaml
type: mqtt
topic: mbmd/charger/maxcurrent
payload: ${var:%d}
```

For write access, the data is provided using the `payload` attribute. If `payload` is missing, the value will be written in default format.

### Script (read/write)

The `script` plugin executes external scripts to read or update data. This plugin is useful to implement any type of external functionality.

Sample read configuration:

```yaml
type: script
cmd: /bin/bash -c "cat /dev/urandom"
timeout: 5s
```

Sample write configuration:

```yaml
type: script
cmd: /home/user/my-script.sh ${enable:%b} # format boolean enable as 0/1
timeout: 5s
```

### HTTP (read/write)

The `http` plugin executes HTTP requests to read or update data. Includes the ability to read and parse JSON using jq-like queries.

Sample read configuration:

```yaml
type: http
uri: https://volkszaehler/api/data/<uuid>.json?from=now
method: GET # default HTTP method
headers:
- content-type: application/json
auth: # basic authorization
  type: basic
  user: foo
  password: bar
insecure: false # set to true to trust self-signed certificates
jq: .data.tuples[0][1] # parse response json
scale: 0.001 # floating point factor applied to result, e.g. for kW to W conversion
```

Sample write configuration:

```yaml
...
body: %v # only applicable for PUT or POST requests
```

### Websocket (read only)

The `websocket` plugin implements a web socket listener. Includes the ability to read and parse JSON using jq-like queries. It can for example be used to receive messages from Volkszähler's push server.

Sample configuration (read only):

```yaml
type: http
uri: ws://<volkszaehler host:port>/socket
jq: .data | select(.uuid=="<uuid>") .tuples[0][1] # parse message json
scale: 0.001 # floating point factor applied to result, e.g. for Wh to kWh conversion
timeout: 30s # error if no update received in 30 seconds
```

### Combined status (read only)

The `combined` status plugin is used to convert a mixed boolean status of plugged/charging into an EVCC-compatible charger status of A..F. It is typically used together with OpenWB MQTT integration.

Sample configuration (read only):

```yaml
type: combined
plugged:
  type: mqtt
  topic: openWB/lp/1/boolPlugStat
charging:
  type: mqtt
  topic: openWB/lp/1/boolChargeStat
```

## API

EVCC provides a REST and MQTT APIs.

### REST API

- `/api/config`: EVCC static configuration
- `/api/state`: EVCC dynamic state
- `/api/mode`: global charge mode, use `/api/mode/<mode>` to modify
- `/api/targetsoc`: global target SoC, use `/api/targetsoc/<soc>` to modify
- `/api/loadpoints/<id>/mode`: loadpoint charge mode, use `/api/loadpoints/<id>/mode/<mode>` to modify
- `/api/loadpoints/<id>/targetsoc`: loadpoint target SoC, use `/api/loadpoints/<id>/targetsoc/<soc>` to modify

### MQTT API

The MQTT API follows the REST API's structure:

- `evcc`: root topic
- `evcc/updated`: timestamp of last update
- `evcc/site`: site dynamic state
- `evcc/site/mode`: global charge mode, write `<mode>` to `/evcc/site/mode/set` to modify
- `evcc/site/targetsoc`: global target SoC, write `<soc>` to `/evcc/site/targetsoc/set` to modify
- `evcc/loadpoints`: number of available loadpoints
- `evcc/loadpoints/<id>`: loadpoint dynamic state
- `evcc/loadpoints/<id>/mode`: loadpoint charge mode, write `<mode>` to `/evcc/loadpoints/<id>/mode/set` to modify
- `evcc/loadpoints/<id>/targetsoc`: loadpoint target SoC, write `<soc>` to `/evcc/loadpoints/<id>/targetsoc/set` to modify

## Background

EVCC is heavily inspired by [OpenWB](1). However, I found OpenWB's architecture slightly intimidating with everything basically global state and heavily relying on shell scripting. On the other side, especially the scripting aspect is one that contributes to [OpenWB's](1) flexibility.

Hence, for a simplified and stricter implementation of an EV charge controller, the design goals for EVCC were:

- typed language with ability for systematic testing - achieved by using [Go](2)
- structured configuration - supports YAML-based [config file](evcc.dist.yaml)
- avoidance of feature bloat, simple and clean UI - utilizes [Bootstrap](3)
- containerized operation beyond Raspberry Pi - provide multi-arch [Docker Image](4)
- support for multiple load points - tbd

[1]: https://github.com/snaptec/openWB
[2]: https://golang.org
[3]: https://getbootstrap.org
[4]: https://hub.docker.com/repository/docker/andig/evcc
[5]: https://github.com/volkszaehler/mbmd
