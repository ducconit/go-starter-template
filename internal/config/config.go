package config

import (
	"core/config"
	"sync"
)

type Config struct {
	ApiConfig *ApiConfig `mapstructure:"api"`
	AppConfig *AppConfig `mapstructure:"app"`
	mu        sync.RWMutex
}

type ApiConfig struct {
	Host             string   `mapstructure:"host"`
	Port             string   `mapstructure:"port"`
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	EnableSwagger    bool     `mapstructure:"enable_swagger"`
}

type AppConfig struct {
	EnableMetrics bool `mapstructure:"enable_metrics"`
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

func App() *AppConfig {
	cfg.mu.RLock()
	defer cfg.mu.RUnlock()
	return cfg.AppConfig
}

func (c *ApiConfig) Address() string {
	return c.Host + ":" + c.Port
}

func EnableMetrics() bool {
	return App().EnableMetrics
}
