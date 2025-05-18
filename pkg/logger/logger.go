package logger

import (
	"log/slog"
	"os"

	"github.com/lokesh-go/go-api-service/pkg/utils"
)

type Logger struct {
	Log *slog.Logger
}

// New creates a new logger
func New(config *Config) *Logger {
	// Initialise slog
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})
	sLogger := slog.New(logHandler).With(utils.HostNameJSONKey, utils.GetHostname(), utils.EnvJSONKey, config.Env, utils.AppJSONKey, config.AppName, utils.AppVersionJSONKey, config.AppVersion)

	// Returns
	return &Logger{Log: sLogger}
}
