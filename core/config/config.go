package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	s *FileStore
)

type FileStore struct {
	*viper.Viper
	opt Option
}

type Option struct {
	EnvPrefix      string
	FilePath       string
	Env            string
	LoadEnv        bool
	OnConfigChange func(e fsnotify.Event)
}

func Load(opt Option) (*FileStore, error) {
	c := viper.New()
	s = &FileStore{Viper: c, opt: opt}

	// Set default values
	s.setDefault()

	// Load main config file
	configDir := filepath.Dir(opt.FilePath)
	configName := strings.TrimSuffix(filepath.Base(opt.FilePath), filepath.Ext(opt.FilePath))
	configType := strings.TrimPrefix(filepath.Ext(opt.FilePath), ".")

	s.AddConfigPath(configDir)
	s.SetConfigName(configName)
	s.SetConfigType(configType)

	// Enable env vars with prefix
	if opt.EnvPrefix != "" {
		s.SetEnvPrefix(opt.EnvPrefix)
	}
	s.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	s.AutomaticEnv()

	// Read main config
	if err := s.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Load .env files if enabled
	if opt.LoadEnv {
		godotenv.Load()

		// Load environment specific .env file if specified
		if opt.Env != "" {
			loadEnvFile(s, fmt.Sprintf(".env.%s", opt.Env))
		}
	}

	s.WatchConfig()
	s.OnConfigChange(func(e fsnotify.Event) {
		if opt.OnConfigChange != nil {
			opt.OnConfigChange(e)
		}
	})

	return s, nil
}

// loadEnvFile loads environment variables from a file with proper prefix handling
func loadEnvFile(v *FileStore, filename string) error {
	envViper := viper.New()
	envViper.SetConfigFile(filename)
	envViper.SetConfigType("env")

	// Only set prefix if it's not empty
	if v.opt.EnvPrefix != "" {
		envViper.SetEnvPrefix(v.opt.EnvPrefix)
	}
	envViper.AutomaticEnv()

	if err := envViper.ReadInConfig(); err != nil {
		// Skip if file doesn't exist
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		}
		return fmt.Errorf("failed to read env file %s: %w", filename, err)
	}

	// Get all keys from env file
	for _, key := range envViper.AllKeys() {
		// Convert key from APP_DB_NAME to db.name format
		targetKey := strings.ToLower(key)
		if v.opt.EnvPrefix != "" && strings.HasPrefix(targetKey, strings.ToLower(v.opt.EnvPrefix)+"_") {
			targetKey = strings.TrimPrefix(targetKey, strings.ToLower(v.opt.EnvPrefix)+"_")
		}
		targetKey = strings.ReplaceAll(targetKey, "_", ".")

		// Set the value in the main config
		v.Set(targetKey, envViper.Get(key))
	}

	return nil
}

func (s *FileStore) setDefault() {
	s.SetDefault("app.env", "development")
	s.SetDefault("app.port", 3000)
	s.SetDefault("app.host", "localhost")
	s.SetDefault("jwt.secret_key", "secret_key")
	s.SetDefault("jwt.access_token_expire", 3600)    // 1 giờ
	s.SetDefault("jwt.refresh_token_expire", 604800) // 7 ngày
	s.SetDefault("jwt.refresh_token_secret_key", "refresh_secret_key")
	s.SetDefault("app.queue_workers", 2)
	s.SetDefault("api.allow_origins", []string{"http://localhost:5173", "http://localhost:3000"})

	// jwt
	s.SetDefault("jwt.nonce_secret_key", "nonce_secret_key")
}
