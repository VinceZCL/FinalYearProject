package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var instance *ConfigStruct

type ConfigStruct struct {
	cfg *viper.Viper `yaml:"-"`
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

	// Security configuration settings
	Security struct {
		// jwt secret key
		Secretkey string
	}
}

func Get() *ConfigStruct {
	if instance != nil {
		return instance
	}

	cfg := viper.New()

	cfg.SetConfigName("config")
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath("./config")
	cfg.AutomaticEnv()
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := cfg.ReadInConfig(); err != nil {
		fmt.Printf("Config | Read Error: %s", err)
	}

	instance = &ConfigStruct{
		cfg: cfg,
	}

	if err := cfg.Unmarshal(instance); err != nil {
		panic(err)
	}
	return instance
}
