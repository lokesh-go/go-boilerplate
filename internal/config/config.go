package config

import (
	"fmt"

	"github.com/lokesh-go/go-api-microservice/pkg/utils"
)

// Methods for config
type Methods interface {
	Get() (c *config)
}

// Initialize the config
func Initialize(env string) (Methods, error) {
	// Read the config file
	data, err := utils.ReadFile(fmt.Sprintf("internal/config/env/%s/config.yaml", env))
	if err != nil {
		return nil, err
	}

	// Bind the config
	var cfg *config
	err = utils.YAMLUnmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	// Returns the config
	return cfg, nil
}

// Get the config
func (c *config) Get() (conf *config) {
	return c
}
