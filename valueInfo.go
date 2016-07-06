package goconfig

import (
	"reflect"
)

// Value generic Config Value type interface
type Value interface {
	Name() string
	Shorthand() *string
	UsageDescription() string

	TypeOfValue() reflect.Type
	HasAValue() bool

	HasADefaultValue() bool
	DefaultValueAsString() *string

	SetDefaultForConfigFile()
	UpdateValueFromConfigFile()
}

// valueInfo common Config value data
type valueInfo struct {
	name             string
	shorthand        *string
	usageDescription string

	typeOfValue reflect.Type
}

// MakeValueInfo make a new ValueInfo structure
func makeValueInfo(name string, shorthand *string, usageDescription string, typeOfValue reflect.Type) *valueInfo {

	return &valueInfo{
		name:             name,
		shorthand:        shorthand,
		usageDescription: usageDescription,

		typeOfValue: typeOfValue,
	}
}

// Name config/param name of value
func (this *valueInfo) Name() string {
	return this.name
}

// Shorthand shorthand param of value, nil if no shorthand available
func (this *valueInfo) Shorthand() *string {
	return this.shorthand
}

// UsageDescription description of usage of this value
func (this *valueInfo) UsageDescription() string {
	return this.usageDescription
}

// TypeOfValue actual type of value
func (this *valueInfo) TypeOfValue() reflect.Type {
	return this.typeOfValue
}
