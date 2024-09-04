package config

import "os"

type Config struct {
    ServerPort string
}

func LoadConfig() *Config {
    return &Config{
        ServerPort: getEnv("SERVER_PORT", "8080"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
