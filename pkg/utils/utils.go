package utils

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Get environment
func GetEnv(envKey string) (env string) {
	return os.Getenv(envKey)
}

// GetHostname
func GetHostname() (hostname string) {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}

	// Returns
	return hostname
}

// IsEmpty check if the string is empty
func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

// Read the file
func ReadFile(filePath string) (data []byte, err error) {
	// Read file
	data, err = os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Returns
	return data, nil
}

// YAML Unmarshal
func YAMLUnmarshal(data []byte, v any) (err error) {
	return yaml.Unmarshal(data, v)
}
