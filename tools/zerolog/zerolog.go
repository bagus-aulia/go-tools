package zerolog

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// LoggerConfig holds the global logger configuration
type LoggerConfig struct {
	Level         string
	EnableConsole bool
	LogFile       string
	TimeFormat    string
	EnableCaller  bool
	EnableStack   bool
}

// DefaultLoggerConfig returns default logger configuration
func DefaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Level:         "INFO",
		EnableConsole: true,
		LogFile:       "",
		TimeFormat:    time.RFC822,
		EnableCaller:  false,
		EnableStack:   true,
	}
}

// ConfigureZerolog configures zerolog with the specified settings
func ConfigureZerolog(level string, enableConsole bool, logFile string) {
	config := LoggerConfig{
		Level:         level,
		EnableConsole: enableConsole,
		LogFile:       logFile,
		TimeFormat:    time.RFC822,
		EnableCaller:  false,
		EnableStack:   true,
	}
	ConfigureZerologWithConfig(config)
}

// ConfigureZerologWithConfig configures zerolog with full configuration
func ConfigureZerologWithConfig(config LoggerConfig) {
	// Setup error stack tracing
	if config.EnableStack {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}

	// Set time format
	zerolog.TimeFieldFormat = config.TimeFormat

	// Enable caller information if needed
	if config.EnableCaller {
		log.Logger = log.With().Caller().Logger()
	}

	var writers []io.Writer

	// Set log level
	level := parseLogLevel(config.Level)
	zerolog.SetGlobalLevel(level)

	// Console output with colors
	if config.EnableConsole {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: config.TimeFormat,
		}
		writers = append(writers, consoleWriter)
	}

	// File output
	if config.LogFile != "" {
		file, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			writers = append(writers, file)
		}
	}

	// Multi-writer setup
	if len(writers) > 1 {
		multiWriter := zerolog.MultiLevelWriter(writers...)
		log.Logger = zerolog.New(multiWriter).With().Timestamp().Logger()
	} else if len(writers) == 1 {
		log.Logger = zerolog.New(writers[0]).With().Timestamp().Logger()
	}
}

// parseLogLevel converts string level to zerolog level
func parseLogLevel(level string) zerolog.Level {
	switch level {
	case "TRACE":
		return zerolog.TraceLevel
	case "DEBUG":
		return zerolog.DebugLevel
	case "INFO":
		return zerolog.InfoLevel
	case "WARN":
		return zerolog.WarnLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "FATAL":
		return zerolog.FatalLevel
	default:
		return zerolog.ErrorLevel
	}
}
