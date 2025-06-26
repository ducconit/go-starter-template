package config

import (
	"path/filepath"
	"time"
)

var basePath string = "."

func Set(key string, value any) {
	s.Set(key, value)
}

func GetString(key string) string {
	return s.GetString(key)
}

func GetFloat64(key string) float64 {
	return s.GetFloat64(key)
}

func GetInt(key string) int {
	return s.GetInt(key)
}

func GetDuration(key string) time.Duration {
	return s.GetDuration(key)
}

func GetStringSlice(key string) []string {
	return s.GetStringSlice(key)
}

func BasePath(paths ...string) string {
	return filepath.Join(basePath, filepath.Join(paths...))
}

func SetBasePath(path string) {
	basePath = path
}

func GetBasePath() string {
	return basePath
}

func GetBool(key string) bool {
	return s.GetBool(key)
}

func Unmarshal(input any) error {
	return s.Unmarshal(input)
}

func UnmarshalKey(key string, input any) error {
	return s.UnmarshalKey(key, input)
}
