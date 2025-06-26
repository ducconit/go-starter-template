package log

import (
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

type fileLogger struct {
	logger *zerolog.Logger
	mu     sync.Mutex
}

func newFileLogger(config Config) (Logger, error) {
	if config.Filename == "" {
		config.Filename = "storage/logs/app.log"
	}

	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(config.Filename), 0755); err != nil {
		return nil, err
	}

	// Setup log rotation
	logRotate := &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,    // megabytes
		MaxBackups: config.MaxBackups, // number of backups
		MaxAge:     config.MaxAge,     // days
		Compress:   config.Compress,   // disabled by default
	}

	// Create multi-writer to output to both file and console in development
	var writers []io.Writer
	writers = append(writers, logRotate)

	// In non-production, also log to console
	if os.Getenv("APP_ENV") != "production" {
		writers = append(writers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02 15:04:05",
		})
	}

	multi := io.MultiWriter(writers...)

	level := getZerologLevel(config.Level)
	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(multi).With().
		Timestamp().
		Logger()

	return &fileLogger{
		logger: &logger,
	}, nil
}

func (l *fileLogger) Debug(msg string, fields ...Field) {
	l.mu.Lock()
	defer l.mu.Unlock()
	e := l.logger.Debug()
	l.addFields(e, fields)
	e.Msg(msg)
}

func (l *fileLogger) Info(msg string, fields ...Field) {
	l.mu.Lock()
	defer l.mu.Unlock()
	e := l.logger.Info()
	l.addFields(e, fields)
	e.Msg(msg)
}

func (l *fileLogger) Warn(msg string, fields ...Field) {
	l.mu.Lock()
	defer l.mu.Unlock()
	e := l.logger.Warn()
	l.addFields(e, fields)
	e.Msg(msg)
}

func (l *fileLogger) Error(msg string, fields ...Field) {
	l.mu.Lock()
	defer l.mu.Unlock()
	e := l.logger.Error()
	l.addFields(e, fields)
	e.Msg(msg)
}

func (l *fileLogger) Fatal(msg string, fields ...Field) {
	l.mu.Lock()
	defer l.mu.Unlock()
	e := l.logger.Fatal()
	l.addFields(e, fields)
	e.Msg(msg)
}

func (l *fileLogger) WithFields(fields ...Field) Logger {
	l.mu.Lock()
	defer l.mu.Unlock()

	e := l.logger.With()
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	newLogger := e.Logger()

	return &fileLogger{
		logger: &newLogger,
		mu:     sync.Mutex{},
	}
}

func (l *fileLogger) Sync() error {
	// Flush any buffered log entries
	// No-op for file logger as lumberjack handles this
	return nil
}

func (l *fileLogger) addFields(e *zerolog.Event, fields []Field) {
	for _, f := range fields {
		e.Interface(f.Key, f.Value)
	}
}
