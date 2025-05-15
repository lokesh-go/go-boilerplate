package logger

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/lokesh-go/go-boilerplate/pkg/utils"
)

// Logger is the interface for logging messages.
type Logger interface {
	Info(msg string, fields ...any)
	Debug(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, err error, fields ...any)
}

type logger struct {
	config *Config
	logger *slog.Logger
}

// New creates a new logger
func New(config *Config) Logger {
	// Initialise slog
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})
	sLogger := slog.New(logHandler).With(utils.HostNameJSONKey, utils.GetHostname(), utils.EnvJSONKey, config.Env, utils.AppJSONKey, config.AppName, utils.AppVersionJSONKey, config.AppVersion)

	// Returns
	return &logger{config: config, logger: sLogger}
}

// Debug
func (l *logger) Debug(msg string, fields ...any) {
	if !l.config.DebugEnabled {
		return
	}
	l.logger.Debug(msg, fields...)
}

// Info
func (l *logger) Info(msg string, fields ...any) {
	l.logger.Info(msg, fields...)
}

// Warn
func (l *logger) Warn(msg string, fields ...any) {
	l.logger.Warn(msg, fields...)
}

// Error
func (l *logger) Error(msg string, err error, fields ...any) {
	// Get err string from error
	var errString string
	if err != nil {
		errString = err.Error()
	}

	// Add additional fields
	if l.config.CallerSkipNo != 0 {
		callerSkipNo = l.config.CallerSkipNo
	}
	additionalFields := map[string]interface{}{
		utils.ErrorJSONKey: errString,
		utils.StackJSONKey: formatStackTrace(callerSkipNo, runtime.Callers(callerSkipNo, make([]uintptr, 32))),
	}
	fields = append(fields, additionalFields)

	l.logger.Error(msg, fields...)
}

func formatStackTrace(callerSkipNo, pcCount int) (stackTrace string) {
	pc := make([]uintptr, pcCount)
	n := runtime.Callers(callerSkipNo, pc)
	frames := runtime.CallersFrames(pc[:n])

	// String builder
	var sb strings.Builder
	for {
		frame, more := frames.Next()
		sb.WriteString(fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}

	// Returns
	return sb.String()
}
