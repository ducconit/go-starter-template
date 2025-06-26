package config

import (
	"core/config"
	"sync"
)

type Config struct {
	ApiConfig *ApiConfig `mapstructure:"api"`

	mu sync.RWMutex
}

type ApiConfig struct {
	Host             string   `mapstructure:"host"`
	Port             string   `mapstructure:"port"`
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	EnableSwagger    bool     `mapstructure:"enable_swagger"`
}

var cfg Config

func Unmarshal() error {
	return config.Unmarshal(&cfg)
}

func Api() *ApiConfig {
	cfg.mu.RLock()
	defer cfg.mu.RUnlock()
	return cfg.ApiConfig
}

func (c *ApiConfig) Address() string {
	return c.Host + ":" + c.Port
}
