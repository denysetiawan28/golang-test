package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ConfigApps(path string) *DefaultConfig {
	//viper.SetConfigFile(path)
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	//viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("failed to read config")
	}

	var conf DefaultConfig

	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println("failed to unmarshall config")
	}
	fmt.Println(conf)

	return &conf
}
