package bootstrap

import (
	"github.com/lokesh-go/go-boilerplate/internal/config"
	"github.com/lokesh-go/go-boilerplate/internal/dal"
	"github.com/lokesh-go/go-boilerplate/pkg/logger"
)

// Initialize application dependencies
func Initialize() (err error) {
	// Get environment
	env := getEnv()

	// Initialize config
	conf, err := config.Initialize(env)
	if err != nil {
		return err
	}

	// Initialize logger
	logger := logger.New(getLoggerConfig(env, conf))
	logger.Info("logger initialised ...")

	// Initialize data access layer
	dal.Initialize()

	// Returns
	return nil
}

func getLoggerConfig(env string, config config.Methods) *logger.Config {
	// Returns
	return &logger.Config{
		Env:          env,
		AppName:      config.Get().App.Name,
		AppVersion:   config.Get().App.Version,
		DebugEnabled: config.Get().Logger.Debug,
		CallerSkipNo: config.Get().Logger.CallerSkipNo,
	}
}
