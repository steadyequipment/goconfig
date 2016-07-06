package goconfig

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"reflect"

	"strconv"
)

const (
	invalidBoolDefault = false
)

type BoolValue struct {
	*valueInfo

	defaultValue *bool
	actualValue  *bool
}

func makeBoolValue(name string, shorthand *string, defaultValue *bool, usageDescription string) *BoolValue {

	var exampleType bool
	valueInfo := MakeValueInfo(name, shorthand, usageDescription, reflect.TypeOf(exampleType))

	result := &BoolValue{
		valueInfo: valueInfo,

		defaultValue: defaultValue,
	}

	result.Setup()

	return result
}

func (this *BoolValue) HasAValue() bool {
	return this.Value() != nil
}

func (this *BoolValue) HasADefaultValue() bool {
	return this.defaultValue != nil && *this.defaultValue != invalidBoolDefault
}

func (this *BoolValue) DefaultValueAsString() *string {
	if this.defaultValue != nil && *this.defaultValue != invalidBoolDefault {
		asString := strconv.FormatBool(*this.defaultValue)
		return &asString
	}

	return nil
}

func (this *BoolValue) Setup() {

	if this.shorthand != nil {
		this.actualValue = pflag.BoolP(this.name,
			*this.shorthand,
			this.useDefaultWithViper(),
			this.usageDescription)
	} else {
		this.actualValue = pflag.Bool(this.name,
			this.useDefaultWithViper(),
			this.usageDescription)

	}
}

func (this *BoolValue) useDefaultWithViper() bool {
	if this.defaultValue != nil {
		return *this.defaultValue
	}

	return invalidBoolDefault
}

func (this *BoolValue) Value() *bool {

	if this.actualValue != nil {
		return this.actualValue
	}

	return nil
}

func (this *BoolValue) SetDefaultForConfigFile() {
	viper.SetDefault(this.name, this.useDefaultWithViper())
}

func (this *BoolValue) UpdateValueFromConfigFile() {
	*this.actualValue = viper.GetBool(this.name)
}

//////////////////////////////////////////////////////////////////

func (this *ConfigValues) makeAndAddBoolValue(name string, shorthand *string, defaultValue *bool, usageDescription string) *BoolValue {

	result := makeBoolValue(name, shorthand, defaultValue, usageDescription)
	this.allValues[name] = result
	return result
}

func (this *ConfigValues) MakeBoolValueWithShorthandAndDefault(name string, shorthand string, defaultValue bool, usageDescription string) *BoolValue {

	return this.makeAndAddBoolValue(name, &shorthand, &defaultValue, usageDescription)
}

func (this *ConfigValues) MakeBoolValueWithShorthand(name string, shorthand string, usageDescription string) *BoolValue {

	return this.makeAndAddBoolValue(name, &shorthand, nil, usageDescription)
}

func (this *ConfigValues) MakeBoolValueWithDefault(name string, defaultValue bool, usageDescription string) *BoolValue {

	return this.makeAndAddBoolValue(name, nil, &defaultValue, usageDescription)
}

func (this *ConfigValues) MakeBoolValue(name string, usageDescription string) *BoolValue {

	return this.makeAndAddBoolValue(name, nil, nil, usageDescription)
}

func (this *ConfigValues) BoolValue(name string) *bool {

	if this.allValues == nil {
		return nil
	}

	configValue, valid := this.allValues[name]
	if !valid || configValue == nil {
		return nil
	}

	boolConfigValue, valid := configValue.(*BoolValue)
	if !valid || boolConfigValue == nil {
		return nil
	}

	return boolConfigValue.Value()
}
