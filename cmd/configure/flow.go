package configure

import (
	"errors"
	"fmt"

	"github.com/evcc-io/evcc/charger"
	"github.com/evcc-io/evcc/meter"
	"github.com/evcc-io/evcc/templates"
	"github.com/evcc-io/evcc/vehicle"
	"gopkg.in/yaml.v3"
)

// let the user choose a device that is set to support guided setup
// these are typically devices that
// - contain multiple usages but have the same parameters like host, port, etc.
// - devices that typically are installed with additional specific devices (e.g. SMA Home Manager with SMA Inverters)
func (c *CmdConfigure) configureDeviceGuidedSetup() {
	var err error

	var values map[string]interface{}
	var deviceCategory DeviceCategory
	var supportedDeviceCategories []DeviceCategory
	var templateItem templates.Template

	deviceItem := device{}

	for ok := true; ok; {
		fmt.Println()

		templateItem, err = c.processDeviceSelection(DeviceCategoryGuidedSetup)
		if err != nil {
			return
		}

		usageChoices := c.paramChoiceValues(templateItem.Params, templates.ParamUsage)
		if len(usageChoices) == 0 {
			panic("ERROR: Device template is missing valid usages!")
		}
		if len(usageChoices) == 0 {
			usageChoices = []string{string(DeviceCategoryGridMeter), string(DeviceCategoryPVMeter), string(DeviceCategoryBatteryMeter)}
		}

		supportedDeviceCategories = []DeviceCategory{}

		for _, usage := range usageChoices {
			switch usage {
			case string(DeviceCategoryGridMeter):
				supportedDeviceCategories = append(supportedDeviceCategories, DeviceCategoryGridMeter)
			case string(DeviceCategoryPVMeter):
				supportedDeviceCategories = append(supportedDeviceCategories, DeviceCategoryPVMeter)
			case string(DeviceCategoryBatteryMeter):
				supportedDeviceCategories = append(supportedDeviceCategories, DeviceCategoryBatteryMeter)
			}
		}

		// we only ask for the configuration for the first usage
		deviceCategory = supportedDeviceCategories[0]

		values = c.processConfig(templateItem.Params, deviceCategory, false)

		deviceItem, err = c.processDeviceValues(values, templateItem, deviceItem, deviceCategory)
		if err != nil {
			if err != c.errDeviceNotValid {
				fmt.Println()
				fmt.Println(err)
			}
			fmt.Println()
			if !c.askConfigFailureNextStep() {
				return
			}
			continue
		}

		break
	}

	c.configuration.AddDevice(deviceItem, deviceCategory)

	if len(supportedDeviceCategories) > 1 {
		for _, additionalCategory := range supportedDeviceCategories[1:] {
			values[templates.ParamUsage] = additionalCategory.String()
			deviceItem, err := c.processDeviceValues(values, templateItem, deviceItem, additionalCategory)
			if err != nil {
				continue
			}

			c.configuration.AddDevice(deviceItem, additionalCategory)
		}
	}

	fmt.Println()
	fmt.Println(templateItem.Description + " " + c.localizedString("Device_Added", nil))

	c.configureLinkedTypes(templateItem)
}

// let the user configure devices that are marked as being linked to a guided device
// e.g. SMA Inverters, Energy Meter with SMA Home Manager
func (c *CmdConfigure) configureLinkedTypes(templateItem templates.Template) {
	linkedTemplates := templateItem.GuidedSetup.Linked

	if linkedTemplates == nil {
		return
	}

	for _, linkedTemplate := range linkedTemplates {
		for ok := true; ok; {
			deviceItem := device{}

			linkedTemplateItem := templates.ByTemplate(linkedTemplate.Template, string(DeviceClassMeter))
			if len(linkedTemplateItem.Params) == 0 || linkedTemplate.Usage == "" {
				return
			}

			category := DeviceCategory(linkedTemplate.Usage)

			fmt.Println()
			if !c.askYesNo(c.localizedString("AddLinkedDeviceInCategory", localizeMap{"Linked": linkedTemplateItem.Description, "Article": DeviceCategories[category].article, "Category": DeviceCategories[category].title})) {
				break
			}

			values := c.processConfig(linkedTemplateItem.Params, category, false)
			deviceItem, err := c.processDeviceValues(values, linkedTemplateItem, deviceItem, category)
			if err != nil {
				if !errors.Is(err, c.errDeviceNotValid) {
					fmt.Println()
					fmt.Println(err)
				}
				fmt.Println()
				if c.askConfigFailureNextStep() {
					continue
				}

			} else {
				c.configuration.AddDevice(deviceItem, category)

				fmt.Println()
				fmt.Println(linkedTemplateItem.Description + " " + c.localizedString("Device_Added", nil))
			}
			break
		}
	}
}

// let the user select and configure a device from a specific category
func (c *CmdConfigure) configureDeviceCategory(deviceCategory DeviceCategory) (device, error) {
	fmt.Println()
	fmt.Printf("- %s %s\n", c.localizedString("Device_Configure", nil), DeviceCategories[deviceCategory].title)

	device := device{
		Name:  DeviceCategories[deviceCategory].defaultName,
		Title: "",
		Yaml:  "",
	}

	deviceDescription := ""

	for ok := true; ok; {
		fmt.Println()

		templateItem, err := c.processDeviceSelection(deviceCategory)
		if err != nil {
			return device, c.errItemNotPresent
		}

		deviceDescription = templateItem.Description
		values := c.processConfig(templateItem.Params, deviceCategory, false)

		device, err = c.processDeviceValues(values, templateItem, device, deviceCategory)
		if err != nil {
			if err != c.errDeviceNotValid {
				fmt.Println()
				fmt.Println(err)
			}
			fmt.Println()
			if !c.askConfigFailureNextStep() {
				return device, err
			}
			continue
		}

		break
	}

	c.configuration.AddDevice(device, deviceCategory)

	deviceTitle := ""
	if device.Title != "" {
		deviceTitle = " " + device.Title
	}

	fmt.Println()
	fmt.Println(deviceDescription + deviceTitle + " " + c.localizedString("Device_Added", nil))

	return device, nil
}

// create a configured device from a template so we can test it
func (c *CmdConfigure) configureDevice(deviceCategory DeviceCategory, device templates.Template, values map[string]interface{}) (interface{}, error) {
	b, err := device.RenderResult(false, values)
	if err != nil {
		return nil, err
	}

	var instance struct {
		Type  string
		Other map[string]interface{} `yaml:",inline"`
	}

	if err := yaml.Unmarshal(b, &instance); err != nil {
		return nil, err
	}

	var v interface{}

	switch DeviceCategories[deviceCategory].class {
	case DeviceClassMeter:
		v, err = meter.NewFromConfig(instance.Type, "", instance.Other)
	case DeviceClassCharger:
		v, err = charger.NewFromConfig(instance.Type, "", instance.Other)
	case DeviceClassVehicle:
		v, err = vehicle.NewFromConfig(instance.Type, "", instance.Other)
	}

	return v, err
}
