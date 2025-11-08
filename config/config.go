package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	ServerPort  int
	Database    DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Load() (*Config, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, fmt.Errorf(".env file not found: %v", err)
	}

	cfg := &Config{
		Environment: strings.ToLower(getStringEnv("ENVIRONMENT")),
		ServerPort:  getIntEnv("SERVER_PORT"),
		Database: DatabaseConfig{
			Host:     getStringEnv("DB_HOST"),
			Port:     getIntEnv("DB_PORT"),
			User:     getStringEnv("DB_USER"),
			Password: getStringEnv("DB_PASSWORD"),
			DBName:   getStringEnv("DB_NAME"),
			SSLMode:  getStringEnv("DB_SSLMODE"),
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func getStringEnv(key string) string {
	return os.Getenv(key)
}

func getIntEnv(key string) int {
	valueStr := os.Getenv(key)

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0
	}

	return value
}

func (c *Config) Validate() error {
	validEnvs := map[string]bool{
		"dev":  true,
		"prod": true,
	}

	if !validEnvs[c.Environment] {
		return fmt.Errorf("invalid environment: %s (must be: dev, prod)", c.Environment)
	}

	if c.ServerPort < 1 || c.ServerPort > 65535 {
		return fmt.Errorf("SERVER_PORT must be between 1 and 65535")
	}

	if c.Database.Host == "" {
		return fmt.Errorf("DB_HOST is required")
	}

	if c.Database.Port < 1 || c.Database.Port > 65535 {
		return fmt.Errorf("DB_PORT must be between 1 and 65535")
	}

	if c.Database.User == "" {
		return fmt.Errorf("DB_USER is required")
	}

	if c.Database.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	if c.Database.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	validSSLModes := map[string]bool{
		"disable":     true,
		"require":     true,
		"verify-ca":   true,
		"verify-full": true,
	}

	if !validSSLModes[c.Database.SSLMode] {
		return fmt.Errorf("invalid DB_SSLMODE: %s", c.Database.SSLMode)
	}

	return nil
}

func (c *Config) IsDevelopment() bool {
	return c.Environment == "dev"
}

func (c *Config) IsProduction() bool {
	return c.Environment == "prod"
}

func (c *Config) GetServerPort() string {
	return ":" + strconv.Itoa(c.ServerPort)
}
