package examples

import (
	"fmt"

	"github.com/GeekSpotlight/go-viper-example/config"
)

func init() {
	config.LoadConfigs(config.App, config.Database, config.External)
}

func RunMultipleConfig() {
	fmt.Println("---- testing multiple configuration ---")
	printConfig(config.App, "application.name")
	printConfig(config.Database, "database.host")
	printConfig(config.External, "external.http.getEmployee.url")
}

func printConfig(configName, key string) {
	fmt.Printf("value of '%s' from config '%s' is '%s'\n", key, configName, config.GetByConfig(configName, key))
}
