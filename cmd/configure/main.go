package configure

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/server"
	"github.com/evcc-io/evcc/templates"
	"github.com/evcc-io/evcc/util"
)

type CmdConfigure struct {
	configuration Configure
	log           *util.Logger
}

// start the interactive configuration
func (c *CmdConfigure) Run(log *util.Logger, logLevel string) {
	c.log = log

	defaultLevel := "error"
	if logLevel != "" {
		defaultLevel = logLevel
	}
	util.LogLevel(defaultLevel, map[string]string{})
	c.log.INFO.Printf("evcc %s (%s)", server.Version, server.Commit)

	fmt.Println()
	fmt.Println("Die nächsten Schritte führen durch die Einrichtung einer Konfigurationsdatei für evcc.")
	fmt.Println("Beachte dass dieser Prozess nicht alle möglichen Szenarien berücksichtigen kann.")
	fmt.Println("Durch Drücken von CTRL-C kann der Prozess abgebrochen werden.")
	fmt.Println()
	fmt.Println("ACHTUNG: Diese Funktionalität hat experimentellen Status!")
	fmt.Println("  D.h. es kann möglich sein, dass die hiermit erstellen Konfigurationsdatei")
	fmt.Println("  in einem Update nicht mehr funktionieren könnten und neu erzeugt werden müsste.")
	fmt.Println("  Wir freuen uns auf euer Feedback auf https://github.com/evcc-io/evcc/discussions/")
	fmt.Println()
	fmt.Println("Auf geht`s:")

	fmt.Println()
	fmt.Println("Wähle eines der folgenden PV Komplettsysteme aus, oder '" + itemNotPresent + "' falls keines dieser Geräte vorhanden ist")
	c.configureDeviceGuidedSetup()

	c.configureDevices(DeviceCategoryGridMeter, false)
	c.configureDevices(DeviceCategoryPVMeter, true)
	c.configureDevices(DeviceCategoryBatteryMeter, true)
	c.configureDevices(DeviceCategoryVehicle, true)
	c.configureLoadpoints()
	c.configureSite()

	yaml, err := c.configuration.RenderConfiguration()
	if err != nil {
		c.log.FATAL.Fatal(err)
	}

	fmt.Println()

	filename := "evcc.yaml"

	for ok := true; ok; {
		_, err := os.Open(filename)
		if errors.Is(err, os.ErrNotExist) {
			break
		}

		fmt.Printf("Die Datei %s existiert bereits.\n", filename)
		if c.askYesNo("Soll die Datei überschrieben werden") {
			break
		}

		filename = c.askValue(question{
			label:        "Gib einen neuen Dateinamen an",
			exampleValue: "evcc_neu.yaml",
			required:     true})
	}

	err = os.WriteFile(filename, yaml, 0755)
	if err != nil {
		fmt.Printf("Die Konfiguration konnte nicht in die Datei %s gespeicher werden", filename)
		c.log.FATAL.Fatal(err)
	}
	fmt.Printf("Deine Konfiguration wurde erfolgreich in die Datei %s gespeichert.\n", filename)
}

// ask device specfic questions
func (c *CmdConfigure) configureDevices(deviceCategory DeviceCategory, askMultiple bool) []device {
	var devices []device

	if deviceCategory == DeviceCategoryGridMeter && c.configuration.MetersOfCategory(deviceCategory) > 0 {
		return nil
	}

	additionalMeter := ""
	if c.configuration.MetersOfCategory(deviceCategory) > 0 {
		additionalMeter = "noch "
	}

	fmt.Println()
	if !c.askYesNo("Möchtest du " + additionalMeter + DeviceCategories[deviceCategory].article + " " + DeviceCategories[deviceCategory].title + " hinzufügen") {
		return nil
	}

	for ok := true; ok; {
		device, err := c.configureDeviceCategory(deviceCategory)
		if err != nil {
			break
		}
		devices = append(devices, device)

		if !askMultiple {
			break
		}

		fmt.Println()
		if !c.askYesNo("Möchtest du noch " + DeviceCategories[deviceCategory].article + " " + DeviceCategories[deviceCategory].title + " hinzufügen") {
			break
		}
	}

	return devices
}

// ask loadpoint specific questions
func (c *CmdConfigure) configureLoadpoints() {
	fmt.Println()
	fmt.Println("- Ladepunkt(e) einrichten")

	for ok := true; ok; {

		loadpointTitle := c.askValue(question{
			label:        "Titel des Ladepunktes",
			defaultValue: defaultTitleLoadpoint,
			required:     true})
		loadpoint := loadpoint{
			Title:      loadpointTitle,
			Phases:     3,
			MinCurrent: 6,
		}

		charger, err := c.configureDeviceCategory(DeviceCategoryCharger)
		if err != nil {
			break
		}

		loadpoint.Charger = charger.Name

		if !charger.ChargerHasMeter {
			if c.askYesNo("Die Wallbox hat keinen Ladestromzähler. Hast du einen externen Zähler dafür installiert der verwendet werden kann") {
				chargeMeter, err := c.configureDeviceCategory(DeviceCategoryChargeMeter)
				if err != nil {
					break
				}

				loadpoint.ChargeMeter = chargeMeter.Name
			}
		}

		vehicles := c.configuration.DevicesOfClass(DeviceClassVehicle)
		if len(vehicles) == 1 {
			loadpoint.Vehicles = append(loadpoint.Vehicles, vehicles[0].Name)
		} else if len(vehicles) > 1 {
			for _, vehicle := range vehicles {
				if c.askYesNo("Wird das Fahrzeug " + vehicle.Title + " hier laden?") {
					loadpoint.Vehicles = append(loadpoint.Vehicles, vehicle.Name)
				}
			}
		}

		powerChoices := []string{"3,6kW", "11kW", "22kW", "Other"}
		fmt.Println()
		powerIndex, _ := c.askChoice("Was ist die maximale Leistung, welche die Wallbox zur Verfügung stellen kann?", powerChoices)
		switch powerIndex {
		case 0:
			loadpoint.MaxCurrent = 16
			loadpoint.Phases = 1
		case 1:
			loadpoint.MaxCurrent = 16
		case 2:
			loadpoint.MaxCurrent = 32
		case 3:
			amperage := c.askValue(question{
				label:     "Was ist die maximale Stromstärke welche die Wallbox auf einer Phase zur Verfügung stellen kann?",
				valueType: templates.ParamValueTypeInt,
				required:  true})
			loadpoint.MaxCurrent, _ = strconv.Atoi(amperage)

			phaseChoices := []string{"1", "2", "3"}
			fmt.Println()
			phaseIndex, _ := c.askChoice("Mit wievielen Phasen ist die Wallbox angeschlossen?", phaseChoices)
			loadpoint.Phases = phaseIndex + 1
		}

		chargingModes := []string{string(api.ModeOff), string(api.ModeNow), string(api.ModeMinPV), string(api.ModePV)}
		ladeModi := []string{"Aus", "Sofort (mit größtmöglicher Leistung)", "Min+PV (mit der kleinstmöglichen Leistung, schneller wenn genügend PV Überschuss vorhanden ist)", "PV (Nur mit PV Überschuß)"}
		fmt.Println()
		modeChoice, _ := c.askChoice("Was sollte der Standard-Lademodus sein, wenn ein Fahrzeug angeschlossen wird?", ladeModi)
		loadpoint.Mode = chargingModes[modeChoice]

		c.configuration.AddLoadpoint(loadpoint)

		fmt.Println()
		if !c.askYesNo("Möchtest du einen weiteren Ladepunkt hinzufügen") {
			break
		}
	}
}

// ask site specific questions
func (c *CmdConfigure) configureSite() {
	fmt.Println()
	fmt.Println("- Richte deinen Standort ein")

	siteTitle := c.askValue(question{
		label:        "Titel des Standortes",
		defaultValue: defaultTitleSite,
		required:     true})
	c.configuration.SetSiteTitle(siteTitle)
}
