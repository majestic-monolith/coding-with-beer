package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	Port        string `mapstructure:"PORT"`
}

func LoadConfig() (Config, error) {
	var cfg Config

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("error reading .env file: %w", err)
	}

	viper.SetDefault("PORT", "8080")

	cfg.DatabaseURL = viper.GetString("DATABASE_URL")

	if cfg.DatabaseURL == "" {
		return cfg, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	return cfg, nil
}
