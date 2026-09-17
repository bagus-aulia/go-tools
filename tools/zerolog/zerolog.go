package zerolog

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LoggerConfig holds the global logger configuration
type LoggerConfig struct {
	Level         string
	EnableConsole bool
	LogFile       string
	TimeFormat    string
	EnableCaller  bool
	EnableStack   bool
	// Log rotation settings
	MaxSize    int  // Maximum file size in MB before rotation (default: 100)
	MaxBackups int  // Maximum number of old log files to keep (default: 3)
	MaxAge     int  // Maximum age of a log file in days before it's deleted (default: 30)
	Compress   bool // Whether to compress rotated log files (default: true)
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
		MaxSize:       100,  // 100 MB
		MaxBackups:    3,    // Keep 3 old files
		MaxAge:        30,   // Delete files older than 30 days
		Compress:      true, // Compress rotated files
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
		MaxSize:       100,
		MaxBackups:    3,
		MaxAge:        30,
		Compress:      true,
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

	// File output with rotation
	if config.LogFile != "" {
		lumberjackLogger := &lumberjack.Logger{
			Filename:   config.LogFile,
			MaxSize:    config.MaxSize,    // megabytes
			MaxBackups: config.MaxBackups, // number of backups
			MaxAge:     config.MaxAge,     // days
			Compress:   config.Compress,   // compress rotated files
		}
		writers = append(writers, lumberjackLogger)
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
