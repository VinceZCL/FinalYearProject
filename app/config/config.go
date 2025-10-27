package config

import (
	"github.com/spf13/viper"
)

var cfg *viper.Viper

type ConfigStruct struct {
	// Database configuration settings.
	Database struct {
		// Database host address.
		Host string
		// Port on which the database is running.
		Port int
		// Name of the database.
		Name string
		// Database username.
		User string
		// Database password.
		Password string
	}
}

func LoadConfig() (*ConfigStruct, error) {
	cfg = viper.New()

	cfg.SetConfigName("config")
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath("./config")

	if err := cfg.ReadInConfig(); err != nil {
		return nil, err
	}

	var config ConfigStruct
	if err := cfg.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func Viper() *viper.Viper {
	return cfg
}
