package configure

import (
	"errors"
	"strings"

	"github.com/evcc-io/evcc/templates"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

func (c *CmdConfigure) setDefaultTexts() {
	c.errItemNotPresent = errors.New(c.localizedString("Error_ItemNotPresent", nil))
	c.errDeviceNotValid = errors.New(c.localizedString("Error_DeviceNotValid", nil))

	c.updateDeviceCategoryTexts(DeviceCategoryCharger, "Category_ChargerTitle", "Category_ChargerArticle")
	c.updateDeviceCategoryTexts(DeviceCategoryGuidedSetup, "Category_SystemTitle", "Category_SystemArticle")
	c.updateDeviceCategoryTexts(DeviceCategoryGridMeter, "Category_GridMeterTitle", "Category_GridMeterArticle")
	c.updateDeviceCategoryTexts(DeviceCategoryPVMeter, "Category_PVMeterTitle", "Category_PVMeterArticle")
	c.updateDeviceCategoryTexts(DeviceCategoryBatteryMeter, "Category_BatteryMeter", "Category_BatteryMeterArticle")
	c.updateDeviceCategoryTexts(DeviceCategoryChargeMeter, "Category_ChargeMeterTitle", "Category_ChargeMeterArticle")
	c.updateDeviceCategoryTexts(DeviceCategoryVehicle, "Category_VehicleTitle", "Category_VehicleArticle")
}

func (c *CmdConfigure) updateDeviceCategoryTexts(category DeviceCategory, title, article string) {
	data := DeviceCategories[category]
	data.title = c.localizedString(title, nil)
	data.article = c.localizedString(article, nil)
	DeviceCategories[category] = data
}

func (c *CmdConfigure) localizedString(key string, templateData localizeMap) string {
	return c.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: templateData,
	})
}

func (c *CmdConfigure) userFriendlyTexts(param templates.Param) templates.Param {
	result := param

	if result.ValueType == "" || (result.ValueType != "" && !funk.ContainsString(templates.ParamValueTypes, result.ValueType)) {
		result.ValueType = templates.ParamValueTypeString
	}

	switch strings.ToLower(result.Name) {
	case "title":
		result.Name = c.localizedString("UserFriendly_Title_Name", nil)
		if result.Help.String(c.lang) == "" {
			result.Help.SetString(c.lang, c.localizedString("UserFriendly_Title_Help", nil))
		}
	case "device":
		result.Name = c.localizedString("UserFriendly_Device_Name", nil)
	case "baudrate":
		result.Name = c.localizedString("UserFriendly_Baudrate_Name", nil)
	case "comset":
		result.Name = c.localizedString("UserFriendly_ComSet_Name", nil)
	case "host":
		result.Name = c.localizedString("UserFriendly_Host_Name", nil)
	case "port":
		result.Name = c.localizedString("UserFriendly_Port_Name", nil)
		result.ValueType = templates.ParamValueTypeNumber
	case "user":
		result.Name = c.localizedString("UserFriendly_User_Name", nil)
	case "password":
		result.Name = c.localizedString("UserFriendly_Password_Name", nil)
	case "capacity":
		result.Name = c.localizedString("UserFriendly_Capacity_Name", nil)
		if result.Example == "" {
			result.Example = "41.5"
		}
		result.ValueType = templates.ParamValueTypeFloat
	case "vin":
		result.Name = c.localizedString("UserFriendly_Vin_Name", nil)
		if result.Help.String(c.lang) == "" {
			result.Help.SetString(c.lang, c.localizedString("UserFriendly_Vin_Help", nil))
		}
	case "identifier":
		result.Name = c.localizedString("UserFriendly_Identifier_Name", nil)
		if result.Help.String(c.lang) == "" {
			result.Help.SetString(c.lang, c.localizedString("UserFriendly_Identifier_Help", nil))
		}
	case "standbypower":
		result.Name = c.localizedString("UserFriendly_StandByPower_Name", nil)
		if result.Help.String(c.lang) == "" {
			result.Help.SetString(c.lang, c.localizedString("UserFriendly_StandByPower_Help", nil))
		}
		result.ValueType = templates.ParamValueTypeNumber
	}
	return result
}
