package main

import (
    "gopkg.in/yaml.v2"
    "log"
    "os"
    "time"
)

// Config holds the configuration values loaded from environment variables and config.yaml
type Config struct {
    ServerPort  string `yaml:"server_port"`
    LogLevel    string `yaml:"log_level"`
    Timeout     time.Duration `yaml:"timeout"`
    FeatureFlags map[string]bool `yaml:"feature_flags"`
}

// LoadConfig loads the configuration from environment variables and a YAML file.
func LoadConfig(filePath string) (*Config, error) {
    config := &Config{
        ServerPort:  getEnv("SERVER_PORT", "8080"),
        LogLevel:    getEnv("LOG_LEVEL", "INFO"),
        Timeout:     time.Duration(getEnvInt("TIMEOUT", 30)) * time.Second,
        FeatureFlags: getFeatureFlags(),
    }

    if err := loadYAMLConfig(filePath, config); err != nil {
        return nil, err
    }

    return config, validateConfig(config)
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
    envValue := getEnv(key, "")
    if envValue == "" {
        return defaultValue
    }
    value, err := strconv.Atoi(envValue)
    if err != nil {
        return defaultValue
    }
    return value
}

func getFeatureFlags() map[string]bool {
    // Load feature flags from environment variables or return defaults
    return map[string]bool{
        "feature_x": getEnv("FEATURE_X", "false") == "true",
        "feature_y": getEnv("FEATURE_Y", "false") == "true",
    }
}

func loadYAMLConfig(filePath string, config *Config) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := yaml.NewDecoder(file)
    return decoder.Decode(config)
}

func validateConfig(config *Config) error {
    if config.ServerPort == "" {
        return fmt.Errorf("missing required field: ServerPort")
    }
    if config.LogLevel == "" {
        return fmt.Errorf("missing required field: LogLevel")
    }
    return nil
}