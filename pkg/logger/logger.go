package logger

import (
	"os"

	"cgoffline/pkg/config"

	"github.com/sirupsen/logrus"
)

// Logger is a global logger instance
var Logger *logrus.Logger

// InitLogger initializes the logger with the given configuration
func InitLogger(cfg config.LoggingConfig) {
	Logger = logrus.New()

	// Set log level
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)

	// Set log format
	switch cfg.Format {
	case "json":
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		})
	default:
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	// Set output
	Logger.SetOutput(os.Stdout)
}

// GetLogger returns the global logger instance
func GetLogger() *logrus.Logger {
	if Logger == nil {
		// Initialize with default config if not already initialized
		InitLogger(config.LoggingConfig{
			Level:  "info",
			Format: "json",
		})
	}
	return Logger
}
