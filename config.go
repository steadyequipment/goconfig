package config

import (
	"github.com/spf13/pflag"

	utility "github.com/steadyequipment/goutility"
)

// TODO: check for duplicate param names or shorthands
// TODO: store order variables were added
type ConfigValues struct {
	HeaderName        *string
	HeaderDescription *string

	configFile *StringValue

	allValues map[string]Value

	requiredValues []string

	FooterName        *string
	FooterDescription *string
}

func MakeConfigValues() *ConfigValues {
	return &ConfigValues{
		allValues: make(map[string]Value, 0),
	}
}

func (this *ConfigValues) enableConfigFileActual(name string, shorthand *string, defaultValue *string, usageDescription string) {

	this.configFile = makeStringValue(name, shorthand, defaultValue, usageDescription)
}

func (this *ConfigValues) EnableConfigFileWithShorthandAndDefault(name string, shorthand string, defaultValue string, usageDescription string) {

	this.enableConfigFileActual(name, &shorthand, &defaultValue, usageDescription)
}

func (this *ConfigValues) EnableConfigFileWithShorthand(name string, shorthand string, usageDescription string) {

	this.enableConfigFileActual(name, &shorthand, nil, usageDescription)
}

func (this *ConfigValues) EnableConfigFileWithDefault(name string, defaultValue string, usageDescription string) {

	this.enableConfigFileActual(name, nil, &defaultValue, usageDescription)
}

func (this *ConfigValues) EnableConfigFile(name string, usageDescription string) {

	this.enableConfigFileActual(name, nil, nil, usageDescription)
}

func (this *ConfigValues) Parse(configFileAsWell bool) (*string, error) {
	pflag.Parse()

	if configFileAsWell {
		return this.ParseFile()
	} else {
		return nil, nil
	}
}

func (this *ConfigValues) ClearValues() {
	this.allValues = nil
}

// TODO: more closely couple required and values so we can't have required values that don't exist
func (this *ConfigValues) AddRequiredValue(requiredValueName string) {
	if this.requiredValues == nil {
		this.requiredValues = make([]string, 0)
	}

	this.requiredValues = append(this.requiredValues, requiredValueName)
}

func (this *ConfigValues) IsRequiredValue(checkName string) bool {

	isRequiredValue := false
	for _, name := range this.requiredValues {
		if name == checkName {
			isRequiredValue = true
			break
		}
	}

	return isRequiredValue
}

func (this *ConfigValues) CheckRequiredValues() error {

	for _, name := range this.requiredValues {

		configValue, valid := this.allValues[name]
		if valid && configValue != nil {

			if configValue.HasAValue() {
				continue
			}

			return utility.NewError("Required: value %s has not been set", name)
		}

		return utility.NewError("Required: value %s has not been set up as a config value", name)
	}

	return nil
}

func (this *ConfigValues) DoViperPflagsCrapMagicParse() (*string, error) {

	configFileName, parseError := this.Parse(true)

	this.SetDefaultsForConfigFile()
	this.UpdateValueFromConfigFile()

	this.Parse(false)

	return configFileName, parseError
}
