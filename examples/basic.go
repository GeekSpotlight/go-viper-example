package examples

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Application struct {
		Name string // to hold application.name
	}
	AppName string // for alias
}

func init() {
	loadConfigurations()
}

func loadConfigurations() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	viper.RegisterAlias("appName", "application.name")
}

func RunBasicExamplesRead() {
	appName := viper.Get("application.name")
	appNameAlias := viper.Get("appName")

	fmt.Println("---- testing config loading ----")
	fmt.Println("application name:", appName)
	fmt.Println("application name alias:", appNameAlias) // prints with alias
}

func RunBasicExamplesUnmarshal() {
	appConfig := getAppConfig()

	fmt.Println("---- testing config unmarshalling ----")
	fmt.Println("application name:", appConfig.Application.Name)
	fmt.Println("application name alias:", appConfig.AppName) // prints with alias
}

func getAppConfig() AppConfig {
	var appConfig AppConfig

	viper.Unmarshal(&appConfig)

	return appConfig
}
