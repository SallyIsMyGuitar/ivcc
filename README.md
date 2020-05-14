# evcc

[![Build Status](https://travis-ci.org/andig/evcc.svg?branch=master)](https://travis-ci.org/andig/evcc)

EVCC is an extensible EV Charge Controller with PV integration implemented in [Go](2).

## Features

- simple and clean user interface
- multiple [chargers](#charger): Wallbe, Phoenix (includes ESL Walli), go-eCharger, NRGKick (Bluetooth and Connect), SimpleEVSE, EVSEWifi, KEBA/BMW, openWB, Mobile Charger Connect, any other charger using scripting
- multiple [meters](#meter): ModBus (Eastron SDM, MPM3PM, SBC ALE3 and many more), SMA Home Manager and SMA Energy Meter, KOSTAL Smart Energy Meter (KSEM, EMxx), any Sunspec-compatible inverter or home battery devices (Fronius, SMA, SolarEdge, KOSTAL, STECA, E3DC), Tesla PowerWall
- different [vehicles](#vehicle) to show battery status: Audi (eTron), BMW (i3), Tesla, Nissan (Leaf), any other vehicle using scripting
- [plugins](#plugins) for integrating with hardware devices and home automation: Modbus (meters and grid inverters), MQTT and shell scripts
- status notifications using [Telegram](https://telegram.org) and [PushOver](https://pushover.net)
- logging using [InfluxDB](https://www.influxdata.com) and [Grafana](https://grafana.com/grafana/)
- soft ramp-up/ramp-down of charge current ensures contactor only switched at minimum current
- electric contactor protection
- REST API

![Screenshot](docs/screenshot.png)

## Index

- [Getting started](#getting-started)
- [Installation](#installation)
- [Configuration](#configuration)
  - [Charger](#charger)
  - [Meter](#meter)
  - [Vehicle](#vehicle)
  - [Loadpoint](#loadpoint)
  - [Considerations](#considerations)
- [Plugins](#plugins)
  - [Modbus](#modbus-read-only)
  - [MQTT](#mqtt-readwrite)
  - [Script](#script-readwrite)
  - [HTTP](#http-readwrite)
  - [Combined status](#combined-status-read-only)
- [Developer information](#developer-information)
- [Background](#background)

## Getting started

1. Install EVCC. For details see [installation](#installation).
2. Copy the default configuration file `evcc.dist.yaml` to `evcc.yaml` and open for editing.
3. To create a minimal setup you need a [meter](#meter) (either grid meter or pv generation meter) and a supported [charger](#charger). A pv meter can also be replaced by a pv inverter. Both need be combined to a [loadpoint](#loadpoint).
4. Configure both meter(s) and charger by:
    - choosing the appropriate `type`
    - add a `name` attribute than can later be referred to
    - add configuration details depending on `type`
  See `evcc.dist.yaml` for examples.
5. Test your meter, charger and optional vehicle configuration by running

        evcc meter|charger|vehicle

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

To build EVCC from source, [Go](2) 1.13 is required:

    make

**Note**: EVCC comes without any guarantee. You are using this software **entirely** at your own risk. It is your responsibility to verify it is working as intended.
EVCC requires a supported charger and a combination of grid, PV and charge meter.
All components **must** be installed by a certified professional.

## Configuration

The EVCC consists of four basic elements: *Charger*, *Meter* and *Vehicle* individually configured and attached to *Loadpoints*.

### Charger

Charger is responsible for handling EV state and adjusting charge current. Available charger implementations are:

- `wallbe`: Wallbe Eco chargers (see [Preparation](#wallbe-preparation)). For older Wallbe boxes (pre 2019) with Phoenix EV-CC-AC1-M3-CBC-RCM-ETH controllers make sure to set `legacy: true` to enable correct current configuration.
- `phoenix-emcp`: chargers with Phoenix EM-CP-PP-ETH controllers like the ESL Walli (Ethernet connection).
- `phoenix-evcc`: chargers with Phoenix EV-CC-AC1-M controllers (ModBus connection)
- `simpleevse`: chargers with SimpleEVSE controllers connected via ModBus (e.g. OpenWB)
- `evsewifi`: chargers with SimpleEVSE controllers using [SimpleEVSE-Wifi](https://github.com/CurtRod/SimpleEVSE-WiFi)
- `nrgkick-bt`: NRGKick chargers with Bluetooth connector (Linux only, not supported on Docker)
- `nrgkick-connect`: NRGKick chargers with Connect module
- `go-e`: go-eCharger chargers
- `keba`: KEBA KeContact P20/P30 and BMW chargers (see [Preparation](#keba-preparation))
- `mcc`: Mobile Charger Connect devices (Audi, Bentley, Porsche)
- `default`: default charger implementation using configurable [plugins](#plugins) for integrating any type of charger

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

If using Docker, make sure that the Docker container can receive UDP messages on port 7090 used by KEBA by using [host networking](https://docs.docker.com/network/host/) in Docker:

```sh
docker run --network=host -p 7070:7070 andig/evcc ...
```

### Meter

Meters provide data about power and energy consumption. Available meter implementations are:

- `modbus`: ModBus meters as supported by [MBMD](https://github.com/volkszaehler/mbmd#supported-devices). Configuration is similar to the [ModBus plugin](#modbus-read-only) where `power` and `energy` specify the MBMD measurement value to use:

  ```yaml
  - name: pv
    type: modbus
    model: sdm
    uri: rs485.fritz.box:23
    id: 2
    power: Power # reading as understood by MBMD, leave empty for power default value
    energy: Export # optional reading for total energy values, specify for charge meter
  ```

- `sma`: SMA Home Manager and SMA Energy Meter. Power reading is configured out of the box but can be customized if necessary. To obtain energy readings define the desired Obis code (Import Energy: "1:1.8.0", Export Energy: "1:2.8.0"):

  ```yaml
  - name: sma-home-manager
    type: sdm
    uri: 192.168.1.4
    power: # leave empty for combined import/export power choose obis
    energy: # leave empty to disable or choose obis 1:1.8.0/1:2.8.0
  ```

- `tesla`: Tesla PowerWall meter. Use `value` to choose meter (grid meter: `site`, pv: `solar`, battery: `battery`)

  ```yaml
  - name: powerwall
    type: tesla
    uri: http://192.168.1.4/api/meters/aggregates
    meter: site # grid meter: `site`, pv: `solar`, battery: `battery`
  ```

  *Note*: this could also be implemented using a `default` meter with the `http` plugin.

- `default`: default meter implementation where meter readings- `power` and `energy` are configured using [plugin](#plugins)

  ```yaml
  - name: vzlogger
    type: default
    power:
      type: http # or any other plugin
      ...
    energy:
      type: http # or any other plugin
      ...
  ```

### Vehicle

Vehicle represents a specific EV vehicle and its battery. If vehicle is configured and assigned to the charger, charge status and remaining charge duration become available in the user interface.

Available vehicle implementations are:

- `audi`: Audi (eTron)
- `bmw`: BMW (i3)
- `nissan`: Nissan (Leaf)
- `tesla`: Tesla (any model)
- `renault`: Renault (Zoe, Kangoo ZE)
- `porsche`: Porsche (Taycan)
- `default`: default vehicle implementation using configurable [plugins](#plugins) for integrating any type of vehicle

### Loadpoint

A loadpoint combines meters, charger and vehicle together and adds optional configuration. A minimal loadpoint configuration needs either pv or grid meter and a charger. More meters can be added as needed:

```yaml
loadpoints:
- name: main # name for logging
  charger: wallbe # charger reference
  vehicle: audi # vehicle reference
  meters:
    grid: sdm630 # grid meter reference
    pv: sma # pv meter reference
```

More options are documented in the `evcc.dist.yaml` sample configuration.

#### Charge modes

The default *charge mode* upon start of EVCC is configured on the loadpoint. Multiple charge modes are supported:

- **Off**: disable the charger, even if car gets connected.
- **Now** (**Sofortladen**): charge immediately with maximum allowed current.
- **Min + PV**: charge immediately with minimum configured current. Additionally use PV if available.
- **PV**: use PV as available. May not charge the car if PV remains dark.

In general, due to the minimum value of 5% for signalling the EV duty cycle, the charger cannot limit the current to below 6A. If the available power calculation demands a limit less than 6A, handling depends on the charge mode. In **PV** mode, the charger will be disabled until available PV power supports charging with at least 6A. In **Min + PV** mode, charging will continue at minimum current of 6A and charge current will be raised as PV power becomes available again.

### Considerations

For intelligent control of PV power usage, EVCC needs to assess how much residual PV power is available at the grid connection point and how much power the charger actually uses. Various methods are implemented to obtain this information, with different degrees of accuracy.

- **PV meter**: Configuring a *PV meter* is the simplest option. *PV meter* measures the PV generation. The charger is allowed to consume:

      Charge Power = PV Meter Power - Residual Power

  The *pv meter* is expected to deliver negative values for export and should not return positive values.

  *Residual Power* is a configurable assumption how much power remaining facilities beside the charger use.

- **Grid meter**: Configuring a *grid meter* is the preferred option. The *grid meter* is expected to be a two-way meter (import+export) and return the current amount of grid export as negative value measured in Watt (W). The charger is then allowed to consume:

      Charge Power = Current Charge Power - Grid Meter Power - Residual Power

  In this setup, *residual power* is used as margin to account for fluctuations in PV production that may be faster than EVCC's control loop.

- **Battery meter**: *battery meter* is used if a home battery is installed and you want charging the EV take priority over charging the home battery. As the home battery would otherwise "grab" all available PV power, this meter measures the home battery charging power.

  With *grid meter* the charger is then allowed to consume:

      Charge Power = Current Charge Power - Grid Meter Power + Battery Meter Power - Residual Power

  or without *grid meter*

      Charge Power = PV Meter Power + Battery Meter Power - Residual Power

  The *battery meter* is expected to deliver negative values when charging and positive values when discharging.

When using a *grid meter* for accurate control of PV utilization, EVCC needs to be able to determine the current charge power. There are two configurations for determining the *current charge power*:

- **Charge meter**: A *charge meter* is often integrated into the charger but can also be installed separately. EVCC expects the *charge meter* to supply *charge power* in Watt (W) and preferably *total energy* in kWh.
If *total energy* is supplied, it can be used to calculate the *charged energy* for the current charging cycle.

- **No charge meter**: If no charge meter is installed, *charge power* is deducted from *charge current* as controlled by the charger. This method is less accurate than using a *charge meter* since the EV may chose to use less power than EVCC has allowed for consumption.
If the charger supplies *total energy* for the charging cycle this value is preferred over the *charge meter*'s value (if present).

## Plugins

Plugins are used to integrate physical devices and external data sources with EVCC. Plugins support both *read* and *write* access. When using plugins for *write* access, the actual data is provided as variable in form of `${var[:format]}`. If `format` is omitted, data is formatted according to the default Go `%v` [format](https://golang.org/pkg/fmt/). The variable is replaced with the actual data before the plugin is executed.

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
- TCP: Sunspec-compatible grid inverters (SMA, SolarEdge, KOSTAL, Fronius, Steca etc)

Use `value` to define the value to read from the device. All values that are supported by [MBMD](https://github.com/volkszaehler/mbmd/blob/master/meters/measurements.go#L28) are pre-configured.

#### Manual configuration

If the Modbus device is not supported by MBMD, the Modbus register can also be manually configured:

```yaml
type: ...
uri/device/id: ...
register:
  address: 40070
  length: 2 # read length in words
  type: holding # holding or input
  decode: int32 # int16|32|64, uint16|32|64, float32|64 and u|int32s
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
scale: 0.001 # floating point factor applied to result, e.g. for kW to W conversion
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

## Developer information

EVCC has the following internal API. The full documentation is available in GoDoc format in https://pkg.go.dev/github.com/andig/evcc/api.

### Charger API

- `Status()`: get charge controller status (`A...F`)
- `Enabled()`: get charger availability
- `Enable(bool)`: set charger availability
- `MaxCurrent(int)`: set maximum allowed charge current in A

Optionally, charger can also provide:

- `CurrentPower()`: power in W (used if charge meter is not present)

### Meter API

- `CurrentPower()`: power in W
- `TotalEnergy()`: energy in kWh (optional)

### Vehicle API

- `Title()`: vehicle name for display in the configuration UI
- `Capacity()`: battery capacity in kWh
- `ChargeState()`: state of charge in %

Optionally, vehicles can also provide:

- `CurrentPower()`: charge power in W (used if charge meter not present)
- `ChargedEnergy()`: charged energy in kWh
- `ChargeDuration()`: charge duration

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
