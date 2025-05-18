package bootstrap

import (
	"github.com/lokesh-go/go-api-microservice/internal/config"
	"github.com/lokesh-go/go-api-microservice/internal/dal"
	"github.com/lokesh-go/go-api-microservice/internal/dependencies"
	"github.com/lokesh-go/go-api-microservice/internal/server"
	"github.com/lokesh-go/go-api-microservice/pkg/logger"
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
	logger.Log.Info("logger initialised ...")

	// Initialize data access layer
	dal, err := dal.Initialize(conf)
	if err != nil {
		return err
	}
	logger.Log.Info("data access layer initialised ...")

	// Starts server
	err = server.Start(&dependencies.Deps{
		Config: conf,
		Logger: logger,
		DAL:    dal,
	})
	if err != nil {
		return err
	}

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
