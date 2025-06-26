package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	AppPort string `mapstructure:"port"`
	Mode    string `mapstructure:"mode"`
}

type DataBaseConfig struct {
	DBHost     string `mapstructure:"host"`
	DBPort     int    `mapstructure:"port"`
	DBUser     string `mapstructure:"user"`
	DBPassword string `mapstructure:"password"`
	DBName     string `mapstructure:"name"`
	DBSSLMode  string `mapstructure:"sslmode"`
}
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	DataBase DataBaseConfig `mapstructure:"database"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("'config.yaml' file not found!")
		} else {
			return nil, fmt.Errorf("fatal error config file: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %w", err)
	}
	return &cfg, nil
}
