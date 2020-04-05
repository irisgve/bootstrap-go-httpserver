package internal

import (
	"flag"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config contains the full config structure of this application
type Config struct {
	AppName string
	Port    string
}

// GetConfig gets the config setup from env or yaml file source
func GetConfig() (*Config, error) {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetTypeByDefaultValue(true)
	v.AutomaticEnv()
	v.SetConfigType("yaml")

	configFile := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	if *configFile == "" {
		return nil, fmt.Errorf("required flag -config not set")
	}

	v.SetConfigFile(*configFile)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error loading config file: %w", err)
	}

	cfg := &Config{
		AppName: v.GetString("app_name"),
		Port:    v.GetString("port"),
	}
	return cfg, nil
}
