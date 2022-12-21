package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	App      string = "app"
	Database string = "database"
	External string = "external"
)

var configMap map[string]*viper.Viper

func init() {
	configMap = make(map[string]*viper.Viper)
}

func LoadConfigs(configs ...string) {
	for _, configName := range configs {
		viperInstance := viper.New()
		viperInstance.SetConfigName(configName)
		viperInstance.SetConfigType("yaml")
		viperInstance.AddConfigPath("./resources")
		err := viperInstance.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %s %w", configName, err))
		}
		configMap[configName] = viperInstance
	}
}

func GetByConfig(configName, key string) interface{} {
	instance, ok := configMap[configName]
	if !ok {
		return nil
	}
	return instance.Get(key)
}
