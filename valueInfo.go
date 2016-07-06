package config

import (
	"reflect"
)

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

type valueInfo struct {
	name             string
	shorthand        *string
	usageDescription string

	typeOfValue reflect.Type
}

func MakeValueInfo(name string, shorthand *string, usageDescription string, typeOfValue reflect.Type) *valueInfo {

	return &valueInfo{
		name:             name,
		shorthand:        shorthand,
		usageDescription: usageDescription,

		typeOfValue: typeOfValue,
	}
}

func (this *valueInfo) Name() string {
	return this.name
}

func (this *valueInfo) Shorthand() *string {
	return this.shorthand
}

func (this *valueInfo) UsageDescription() string {
	return this.usageDescription
}

func (this *valueInfo) TypeOfValue() reflect.Type {
	return this.typeOfValue
}
