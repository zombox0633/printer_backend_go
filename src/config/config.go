package config

import "os"

type ConfigType struct {
	Port        string
	Environment string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func LoadConfig() ConfigType {
	return ConfigType{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}
