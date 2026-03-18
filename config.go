package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"log"
)

// Config holds the application configuration
type Config struct {
	ServerPort  string `yaml:"server_port"`  // Port for the server to listen on
	LogLevel    string `yaml:"log_level"`    // Log level (e.g., DEBUG, INFO, WARN, ERROR)
	Timeout     int    `yaml:"timeout"`      // Timeout value in seconds
	FeatureFlag bool   `yaml:"feature_flag"` // A feature toggle
}

// LoadConfig reads configuration from the environment and a YAML file
func LoadConfig() (*Config, error) {
	var config Config
	
	// Load from YAML file
	file, err := os.Open("./config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	
	// Load from environment variables
	if port, exists := os.LookupEnv("SERVER_PORT"); exists {
		config.ServerPort = port
	}
	if logLevel, exists := os.LookupEnv("LOG_LEVEL"); exists {
		config.LogLevel = logLevel
	}
	if timeout, exists := os.LookupEnv("TIMEOUT"); exists {
		config.Timeout = timeout
	}
	if featureFlag, exists := os.LookupEnv("FEATURE_FLAG"); exists {
		config.FeatureFlag = featureFlag == "true"
	}

	// Validate required fields
	if config.ServerPort == "" {
		log.Fatal("SERVER_PORT is required")
	}
	if config.LogLevel == "" {
		log.Fatal("LOG_LEVEL is required")
	}
	
	return &config, nil
}