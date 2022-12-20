package main

import (
	"fmt"
	"os"

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

func main() {
	// appName := viper.Get("application.name")
	// appNameAlias := viper.Get("appName")		// prints with alias
	appConfig := getAppConfig()
	fmt.Println("application name:", appConfig.Application.Name)
	fmt.Println("application name alias:", appConfig.AppName) // prints with alias

	singleParam()
	doubleParam()
	multiParam()
	automatic()
}

func singleParam() {
	// setting up env variables
	os.Setenv("TGC_DATAPREFIX", "data with prefix TGC")
	os.Setenv("DATA_NO_PREFIX", "got no prefix")

	vPrefix := viper.New()
	vPrefix.SetEnvPrefix("TGC")
	vPrefix.BindEnv("dataprefix")

	vNoPrefix := viper.New()
	vNoPrefix.BindEnv("data_no_prefix")

	fmt.Println("output 1:", vPrefix.Get("dataprefix"))
	fmt.Println("output 2:", vNoPrefix.Get("data_no_prefix"))
}

func doubleParam() {
	// setting up env variables
	os.Setenv("TGC_DATA", "data with prefix TGC")
	os.Setenv("env_data", "got no prefix")

	vPrefix := viper.New()
	vPrefix.SetEnvPrefix("TGC")
	vPrefix.BindEnv("data", "env_data")        // does not append prefix and is not in full uppercase.
	vPrefix.BindEnv("data_prefix", "TGC_DATA") // explicitly mention full name of the param. prefix is ignored.

	fmt.Println("output 1:", vPrefix.Get("data"))
	fmt.Println("output 2:", vPrefix.Get("data_prefix"))
}

func multiParam() {
	// setting up env variables
	os.Setenv("high_precedence", "precedence level 1")
	os.Setenv("normal_precedence", "precedence level 2")
	os.Setenv("low_precedence", "precedence level 3")
	v := viper.New()

	v.BindEnv("precedence", "no_env_variable_configured", "high_precedence", "normal_precedence", "low_precedence")

	fmt.Println("output:", v.Get("precedence"))
}

func automatic() {
	// setting up env variables
	os.Setenv("TGC_AUTO_1", "an automatic env")
	os.Setenv("TGC_AUTO_2", "another automatic env")

	vPrefix := viper.New()
	vPrefix.SetEnvPrefix("TGC")
	vPrefix.AutomaticEnv()

	vNoPrefix := viper.New()
	vNoPrefix.AutomaticEnv()

	fmt.Println("output 1: ", vPrefix.Get("auto_1"))
	fmt.Println("output 2: ", vPrefix.Get("AUTO_1"))
	fmt.Println("output 3: ", vPrefix.Get("auto_2"))

	// when not using prefix, Get by full name
	fmt.Println("output 4: ", vNoPrefix.Get("TGC_AUTO_1"))
	fmt.Println("output 5: ", vNoPrefix.Get("tgc_auto_1"))
	fmt.Println("output 6: ", vNoPrefix.Get("TGC_AUTO_2"))
}

func getAppConfig() AppConfig {
	var appConfig AppConfig

	viper.Unmarshal(&appConfig)

	return appConfig
}
