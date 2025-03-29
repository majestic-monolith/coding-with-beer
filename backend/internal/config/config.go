package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	ApiSecret         string `mapstructure:"API_SECRET"`
	DatabaseURL       string `mapstructure:"DATABASE_URL"`
	Port              string `mapstructure:"PORT"`
	TokenHourLifespan string `mapstructure:"TOKEN_HOUR_LIFESPAN"`
}

var lock = &sync.Mutex{}

var cfg *Config

func LoadConfig() (Config, error) {
	if cfg != nil {
		return *cfg, nil
	}

	lock.Lock()
	defer lock.Unlock()

	var cfg Config

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("error reading .env file: %w", err)
	}

	viper.SetDefault("PORT", "8080")

	cfg.DatabaseURL = viper.GetString("DATABASE_URL")
	cfg.ApiSecret = viper.GetString("API_SECRET")
	cfg.TokenHourLifespan = viper.GetString("TOKEN_HOUR_LIFESPAN")

	if cfg.DatabaseURL == "" {
		return cfg, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	if cfg.ApiSecret == "" {
		return cfg, fmt.Errorf("API_SECRET environment variable is not set")
	}

	if cfg.TokenHourLifespan == "" {
		return cfg, fmt.Errorf("TOKEN_HOUR_LIFESPAN environment variable is not set")
	}

	return cfg, nil
}
