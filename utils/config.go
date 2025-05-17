package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config stores application configuration
type Config struct {
	ServerPort      int
	AIServiceType   string
	DefaultLanguage string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() Config {
	godotenv.Load()

	config := Config{
		ServerPort:      8080, // Default port
		AIServiceType:   "gemini",
		DefaultLanguage: "en",
	}

	// Override with environment variables if set
	if port := os.Getenv("SERVER_PORT"); port != "" {
		if portNum, err := strconv.Atoi(port); err == nil {
			config.ServerPort = portNum
		}
	}

	if serviceType := os.Getenv("AI_SERVICE_TYPE"); serviceType != "" {
		config.AIServiceType = serviceType
	}

	if language := os.Getenv("DEFAULT_LANGUAGE"); language != "" {
		config.DefaultLanguage = language
	}

	return config
}

// GetServerAddress returns the formatted server address
func (c Config) GetServerAddress() string {
	return fmt.Sprintf(":%d", c.ServerPort)
}
