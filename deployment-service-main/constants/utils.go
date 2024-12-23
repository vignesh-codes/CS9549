package constants

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}

func GetEnvFormattedString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = fmt.Sprintf(defaultValue, ENV)
	}
	return value
}

func GetEnvBool(key string, defaultValue bool) bool {
	envValue := os.Getenv(key)
	if len(envValue) == 0 {
		value := defaultValue
		return value
	}
	value, err := strconv.ParseBool(envValue)
	if err != nil {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	envValue := os.Getenv(key)
	if len(envValue) == 0 {
		value := defaultValue
		return value
	}
	value, err := strconv.Atoi(envValue)
	if err != nil {
		return defaultValue
	}
	return value
}

func GetEnvFloat(key string, defaultValue float64) float64 {
	envValue := os.Getenv(key)
	if len(envValue) == 0 {
		value := defaultValue
		return value
	}
	value, err := strconv.ParseFloat(envValue, 64)
	if err != nil {
		return defaultValue
	}
	return value
}
