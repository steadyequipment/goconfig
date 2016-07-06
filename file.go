package config

import (
	"github.com/spf13/viper"

	"path"
	"path/filepath"

	utility "github.com/steadyequipment/goutility"
)

func (this *ConfigValues) ParseFile() (*string, error) {

	if this.configFile == nil {
		return nil, nil
	}

	configFile := this.configFile.Value()
	if configFile == nil {
		return nil, nil
	}

	configFileLocation, locationError := this.ConfigFileLocation()
	if locationError != nil {
		return configFile, locationError
	}
	if configFileLocation == nil {
		return configFile, utility.NewError("Unable to resolve config file's location")
	}

	configFileAbs, absError := this.ConfigFileAbsLocation()
	if absError != nil {
		return configFileLocation, absError
	}
	if configFileAbs == nil {
		return configFileLocation, utility.NewError("Unable to resolve config file '%s''s absolute location", *configFileLocation)
	}

	configFileName := path.Base(*configFileAbs)
	configPath := (*configFileAbs)[0 : len(*configFileAbs)-len(configFileName)]

	configExtension := path.Ext(configFileName)
	if configExtension[0] == '.' {
		configExtension = configExtension[1:]
	}

	configFileNameWithoutExtension := configFileName[0 : len(configFileName)-(len(configExtension)+1)] // +1 for the '.'
	// TODO:
	// 			fmt.Printf("DEBUG: Using config file %s in path %s\n", configFileName, configPath)

	viper.SetConfigName(configFileNameWithoutExtension)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configExtension)

	readError := viper.ReadInConfig()
	return configFileAbs, readError
}

func (this *ConfigValues) ConfigFileLocation() (*string, error) {

	if this.configFile == nil {
		return nil, utility.NewError("No config file specified")
	}

	configFileLocation := this.configFile.Value()
	if configFileLocation == nil || len(*configFileLocation) <= 0 {
		return nil, utility.NewError("Invalid config file location provided")
	}

	return configFileLocation, nil
}

func (this *ConfigValues) ConfigFileAbsLocation() (*string, error) {

	configFileLocation, locationError := this.ConfigFileLocation()
	if locationError != nil {
		return nil, locationError
	}

	var absError error
	configFileAbs, absError := filepath.Abs(*configFileLocation)
	if absError != nil {
		return nil, absError
	}

	return &configFileAbs, nil
}

func (this *ConfigValues) SetDefaultsForConfigFile() {

	for _, configValue := range this.allValues {
		configValue.SetDefaultForConfigFile()
	}
}

func (this *ConfigValues) UpdateValueFromConfigFile() {

	for _, configValue := range this.allValues {
		configValue.UpdateValueFromConfigFile()
	}
}
