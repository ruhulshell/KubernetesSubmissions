package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}

func GetInt(key string, fallback int) int {
	valueStr := GetString(key, "")
	if valueStr == "" {
		return fallback
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return fallback
	}
	return value
}
