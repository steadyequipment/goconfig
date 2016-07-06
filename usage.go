package config

import (
	"github.com/spf13/pflag"

	"fmt"
	"github.com/fatih/color"

	utility "github.com/steadyequipment/goutility"
)

var (
	sectionNameColor         = color.New(color.FgHiWhite, color.Bold, color.Underline)
	optionColor              = color.New(color.FgWhite, color.Bold)
	optionTypeColor          = color.New(color.Underline)
	requiredDescriptionColor = color.New(color.Bold)
	defaultColor             = color.New(color.Bold)
)

func (this *ConfigValues) SetUsage(usage func()) {
	pflag.Usage = usage
}

func (this *ConfigValues) SetHeaderName(name string) {
	this.HeaderName = &name
}

func (this *ConfigValues) SetHeaderDescription(description string) {
	this.HeaderDescription = &description
}

func (this *ConfigValues) printUsageHeader() {
	if this.HeaderName != nil {
		fmt.Printf("\n")
		fmt.Printf("%s\n", sectionNameColor.SprintFunc()(*this.HeaderName))
	}

	if this.HeaderDescription != nil {
		fmt.Printf("\n")
		fmt.Printf("  %s\n", *this.HeaderDescription)
	}
}

func (this *ConfigValues) SetFooterName(name string) {
	this.FooterName = &name
}

func (this *ConfigValues) SetFooterDescription(description string) {
	this.FooterDescription = &description
}

func (this *ConfigValues) printUsageFooter() {
	if this.FooterName != nil {
		fmt.Printf("\n")
		fmt.Printf("%s\n", sectionNameColor.SprintFunc()(*this.FooterName))
	}

	if this.FooterDescription != nil {
		fmt.Printf("\n")
		fmt.Printf("  %s\n", *this.FooterDescription)
	}
}

func (this *ConfigValues) usageShorthandString(shorthand string) string {
	return optionColor.SprintFunc()("-" + shorthand)
}

func (this *ConfigValues) getNameLength(value Value) int {
	return len("--") + len(value.Name()) + len(" ") + len(value.TypeOfValue().String())
}

func (this *ConfigValues) getLongestName() int {
	result := 0

	for _, value := range this.allValues {
		nameLength := this.getNameLength(value)
		if nameLength > result {
			result = nameLength
		}
	}

	return result
}

func (this *ConfigValues) plainUsageNameTypeString(value Value) string {
	return "--" + value.Name() + " " + value.TypeOfValue().String()
}

func (this *ConfigValues) usageNameTypeString(value Value) string {
	return optionColor.SprintFunc()("--"+value.Name()) + " " + optionTypeColor.SprintFunc()(value.TypeOfValue().String())
}

func (this *ConfigValues) printOptions() {
	// TODO: out, err := exec.Command("stty", "size").Output() for Description wrapping alignment

	leadSpacing := "  "
	shorthandFormat := "%s, "
	shorthandLength := len(fmt.Sprintf(shorthandFormat, "-a"))
	shorthandMissingFormat := utility.StringOfStringRepeated(" ", shorthandLength)

	longestNameLength := this.getLongestName()

	nameFormat := "%s"

	descriptionLeadSpacing := "  "

	for _, value := range this.allValues {
		fmt.Printf(leadSpacing)

		shorthand := value.Shorthand()
		if shorthand != nil {
			fmt.Printf(shorthandFormat, this.usageShorthandString(*shorthand))
		} else {
			fmt.Printf(shorthandMissingFormat)
		}

		usageNameString := this.usageNameTypeString(value)
		plainUsageNameTypeString := this.plainUsageNameTypeString(value)

		padding := utility.StringOfStringRepeated(" ", longestNameLength-len(plainUsageNameTypeString))
		fmt.Printf(nameFormat, usageNameString+padding)

		fmt.Printf(descriptionLeadSpacing)
		if this.IsRequiredValue(value.Name()) {
			fmt.Printf("%s: ", requiredDescriptionColor.SprintFunc()("Required"))
		}
		fmt.Printf(value.UsageDescription())

		if value.HasADefaultValue() {
			defaultValueString := value.DefaultValueAsString()
			if defaultValueString != nil {
				fmt.Printf("      %s: %s", defaultColor.SprintFunc()("Default"), *defaultValueString)
			}
		}

		fmt.Printf("\n")
	}
}

func (this *ConfigValues) PrintUsage() {

	this.printUsageHeader()

	fmt.Printf("\n")
	fmt.Printf("%s\n", sectionNameColor.SprintFunc()("Options"))
	fmt.Printf("\n")
	this.printOptions()

	this.printUsageFooter()

	fmt.Printf("\n")
}

func (this *ConfigValues) printPrompt(prompt string) {
	fmt.Printf("\n")
	fmt.Printf("  %s\n", prompt)
}

func (this *ConfigValues) PrintUsageWithPrompt(prompt string) {
	this.printPrompt(prompt)

	this.PrintUsage()
}
