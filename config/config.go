package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the configuration structure for the application.
type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

// DatabaseConfig represents the database configuration details.
type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

// LoadConfig loads the configuration from a YAML file.
func LoadConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Failed to close file: %v", err)
		}
	}()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
