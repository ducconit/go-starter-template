package log

import "fmt"

// Logger is the interface for logging
// It supports multiple backends (console, file) with thread-safe operations
type Logger interface {
	// Debug logs a message at Debug level
	Debug(msg string, fields ...Field)

	// Info logs a message at Info level
	Info(msg string, fields ...Field)

	// Warn logs a message at Warn level
	Warn(msg string, fields ...Field)

	// Error logs a message at Error level
	Error(msg string, fields ...Field)
	
	// Fatal logs a message at Fatal level and exits the program
	Fatal(msg string, fields ...Field)

	// WithFields returns a new logger with the specified fields
	WithFields(fields ...Field) Logger

	// Sync flushes any buffered log entries
	Sync() error
}

// Field represents a key-value pair for structured logging
type Field struct {
	Key   string
	Value any
}

// NewField creates a new Field
func NewField(key string, value any) Field {
	return Field{Key: key, Value: value}
}

// LoggerType defines the type of logger to use
type LoggerType string

const (
	// ConsoleLogger outputs logs to console (default)
	ConsoleLogger LoggerType = "console"
	// FileLogger outputs logs to files
	FileLogger LoggerType = "file"
)

// Config holds configuration for the logger
type Config struct {
	Type       string // console or file
	Level      string // debug, info, warn, error, fatal
	JSONFormat bool   // output in JSON format
	// FileLogger specific
	Filename   string // log file path
	MaxSize    int    // maximum size in megabytes before rotation (file logger only)
	MaxBackups int    // maximum number of old log files to retain (file logger only)
	MaxAge     int    // maximum number of days to retain old log files (file logger only)
	Compress   bool   // whether to compress rotated log files (file logger only)
}

// NewLogger creates a new logger based on configuration
func NewLogger(config Config) (Logger, error) {
	switch config.Type {
	case "file":
		return newFileLogger(config)
	case "console", "": // Default to console if empty
		return newConsoleLogger(config), nil
	default:
		return nil, fmt.Errorf("unsupported logger type: %s. Use 'console' or 'file'", config.Type)
	}
}
