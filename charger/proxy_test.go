package charger

import (
	"testing"

	"github.com/evcc-io/evcc/templates"
	"github.com/evcc-io/evcc/util/test"
	"github.com/thoas/go-funk"
)

func TestProxyChargers(t *testing.T) {
	test.SkipCI(t)

	for _, tmpl := range templates.ByClass(templates.Charger) {
		tmpl := tmpl

		values := tmpl.Defaults(true)

		// Modbus default test values
		if values[templates.ParamModbus] != nil {
			modbusChoices := tmpl.ModbusChoices()
			if funk.ContainsString(modbusChoices, templates.ModbusChoiceTCPIP) {
				values[templates.ModbusKeyTCPIP] = true
			} else {
				values[templates.ModbusKeyRS485TCPIP] = true
			}
		}

		t.Run(tmpl.Type(), func(t *testing.T) {
			t.Parallel()

			b, err := tmpl.RenderResult(true, values)
			if err != nil {
				t.Logf("%s: %s", tmpl.Template, b)
				t.Error(err)
			}

			_, err = NewFromConfig(tmpl.Type(), values)
			if err != nil && !test.Acceptable(err, acceptable) {
				t.Logf("%s", tmpl.Template)
				t.Error(err)
			}
		})
	}
}
