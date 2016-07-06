package goconfig

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"reflect"

	"strconv"
)

const (
	maxUint           = ^uint(0)
	invalidIntDefault = int(maxUint >> 1)
)

// IntValue int Config Value type
type IntValue struct {
	*valueInfo

	defaultValue *int
	actualValue  *int
}

func makeIntValue(name string, shorthand *string, defaultValue *int, usageDescription string) *IntValue {

	var exampleType int
	valueInfo := makeValueInfo(name, shorthand, usageDescription, reflect.TypeOf(exampleType))

	result := &IntValue{
		valueInfo: valueInfo,

		defaultValue: defaultValue,
	}

	result.Setup()

	return result
}

func (this *IntValue) HasAValue() bool {
	return this.Value() != nil
}

func (this *IntValue) HasADefaultValue() bool {
	return this.defaultValue != nil && *this.defaultValue != invalidIntDefault
}

func (this *IntValue) DefaultValueAsString() *string {
	if this.defaultValue != nil && *this.defaultValue != invalidIntDefault {
		asString := strconv.Itoa(*this.defaultValue)
		return &asString
	}

	return nil
}

func (this *IntValue) Setup() {

	if this.shorthand != nil {
		this.actualValue = pflag.IntP(this.name,
			*this.shorthand,
			this.useDefaultWithViper(),
			this.usageDescription)
	} else {
		this.actualValue = pflag.Int(this.name,
			this.useDefaultWithViper(),
			this.usageDescription)

	}
}

func (this *IntValue) useDefaultWithViper() int {
	if this.defaultValue != nil {
		return *this.defaultValue
	}

	return invalidIntDefault
}

func (this *IntValue) Value() *int {

	if this.actualValue != nil && *this.actualValue != invalidIntDefault {
		return this.actualValue
	}

	return nil
}

func (this *IntValue) SetDefaultForConfigFile() {
	viper.SetDefault(this.name, this.useDefaultWithViper())
}

func (this *IntValue) UpdateValueFromConfigFile() {
	*this.actualValue = viper.GetInt(this.name)
}

//////////////////////////////////////////////////////////////////

func (this *ConfigValues) makeAndAddIntValue(name string, shorthand *string, defaultValue *int, usageDescription string) *IntValue {

	result := makeIntValue(name, shorthand, defaultValue, usageDescription)
	this.allValues[name] = result
	return result
}

func (this *ConfigValues) MakeIntValueWithShorthandAndDefault(name string, shorthand string, defaultValue int, usageDescription string) *IntValue {

	return this.makeAndAddIntValue(name, &shorthand, &defaultValue, usageDescription)
}

func (this *ConfigValues) MakeIntValueWithShorthand(name string, shorthand string, usageDescription string) *IntValue {

	return this.makeAndAddIntValue(name, &shorthand, nil, usageDescription)
}

func (this *ConfigValues) MakeIntValueWithDefault(name string, defaultValue int, usageDescription string) *IntValue {

	return this.makeAndAddIntValue(name, nil, &defaultValue, usageDescription)
}

func (this *ConfigValues) MakeIntValue(name string, usageDescription string) *IntValue {

	return this.makeAndAddIntValue(name, nil, nil, usageDescription)
}

func (this *ConfigValues) IntValue(name string) *int {

	if this.allValues == nil {
		return nil
	}

	configValue, valid := this.allValues[name]
	if !valid || configValue == nil {
		return nil
	}

	intConfigValue, valid := configValue.(*IntValue)
	if !valid || intConfigValue == nil {
		return nil
	}

	return intConfigValue.Value()
}
