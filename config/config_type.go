package config

import "time"

type DefaultConfig struct {
	Apps       Apps       `mapstructure:"apps"`
	Server     Server     `mapstructure:"server"`
	Database   Datasource `mapstructure:"database"`
	//AuthConfig AuthConfig `mapstructure:"authConfig"`
}

type Apps struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Datasource struct {
	Url               string        `mapstructure:"url"`
	Port              string        `mapstructure:"port"`
	DatabaseName      string        `mapstructure:"databaseName"`
	Username          string        `mapstructure:"username"`
	Password          string        `mapstructure:"password"`
	Schema            string        `mapstructure:"schema"`
	ConnectionTimeout time.Duration `mapstructure:"connectionTimeout"`
	MaxIdleConnection int           `mapstructure:"maxIdleConnection"`
	MaxOpenConnection int           `mapstructure:"maxOpenConnection"`
	DebugMode         bool          `mapstructure:"debugMode"`
}

type AuthConfig struct {
	Salt               string `mapstructure:"salt"`
	Key                string `mapstructure:"key"`
	SessionMin         string `mapstructure:"sessionMin"`
	SessionRememberMe  string `mapstructure:"sessionRememberMe"`
	MaximumLoginFailed int    `mapstructure:"maximumLoginFailed"`
}
