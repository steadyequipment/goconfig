package goconfig

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"reflect"
)

const (
	invalidStringDefault = ""
)

type StringValue struct {
	*valueInfo

	defaultValue *string
	actualValue  *string
}

func makeStringValue(name string, shorthand *string, defaultValue *string, usageDescription string) *StringValue {

	var exampleType string
	valueInfo := makeValueInfo(name, shorthand, usageDescription, reflect.TypeOf(exampleType))

	result := &StringValue{
		valueInfo: valueInfo,

		defaultValue: defaultValue,
	}

	result.Setup()

	return result
}

func (this *StringValue) HasAValue() bool {
	return this.Value() != nil
}

func (this *StringValue) HasADefaultValue() bool {
	return this.defaultValue != nil && *this.defaultValue != invalidStringDefault
}

func (this *StringValue) DefaultValueAsString() *string {
	if this.defaultValue != nil && *this.defaultValue != invalidStringDefault {
		return this.defaultValue
	}

	return nil
}

func (this *StringValue) Setup() {

	if this.shorthand != nil {
		this.actualValue = pflag.StringP(this.name,
			*this.shorthand,
			this.useDefaultWithViper(),
			this.usageDescription)
	} else {
		this.actualValue = pflag.String(this.name,
			this.useDefaultWithViper(),
			this.usageDescription)
	}
}

func (this *StringValue) useDefaultWithViper() string {
	if this.defaultValue != nil {
		return *this.defaultValue
	}

	return invalidStringDefault
}

func (this *StringValue) Value() *string {
	if this.actualValue != nil && *this.actualValue != invalidStringDefault {
		return this.actualValue
	}

	return nil
}

func (this *StringValue) SetDefaultForConfigFile() {
	viper.SetDefault(this.name, this.useDefaultWithViper())
}

func (this *StringValue) UpdateValueFromConfigFile() {
	*this.actualValue = viper.GetString(this.name)
}

///////////////////////////////////////////////////////////////

func (this *ConfigValues) makeStringValueActual(name string, shorthand *string, defaultValue *string, usageDescription string) *StringValue {

	result := makeStringValue(name, shorthand, defaultValue, usageDescription)
	this.allValues[name] = result
	return result
}

func (this *ConfigValues) MakeStringValueWithShorthandAndDefault(name string, shorthand string, defaultValue string, usageDescription string) *StringValue {

	return this.makeStringValueActual(name, &shorthand, &defaultValue, usageDescription)
}

func (this *ConfigValues) MakeStringValueWithShorthand(name string, shorthand string, usageDescription string) *StringValue {

	return this.makeStringValueActual(name, &shorthand, nil, usageDescription)
}

func (this *ConfigValues) MakeStringValueWithDefault(name string, defaultValue string, usageDescription string) *StringValue {

	return this.makeStringValueActual(name, nil, &defaultValue, usageDescription)
}

func (this *ConfigValues) MakeStringValue(name string, usageDescription string) *StringValue {

	return this.makeStringValueActual(name, nil, nil, usageDescription)
}

func (this *ConfigValues) StringValue(name string) *string {

	if this.allValues == nil {
		return nil
	}

	configValue, valid := this.allValues[name]
	if !valid || configValue == nil {
		return nil
	}

	stringConfigValue, valid := configValue.(*StringValue)
	if !valid || stringConfigValue == nil {
		return nil
	}

	return stringConfigValue.Value()
}
