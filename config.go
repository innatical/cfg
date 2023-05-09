package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.SetTypeByDefaultValue(true)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("No config file found: Run 'inncfg init' to create one.\n")
	}
}

func configExists() bool {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.SetTypeByDefaultValue(true)

	err := viper.ReadInConfig()

	if err != nil {
		return false
	}

	return true
}
