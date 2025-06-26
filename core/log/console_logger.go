package log

import (
	"io"
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog"
)

type consoleLogger struct {
	logger *zerolog.Logger
	mu     sync.Mutex
}

func newConsoleLogger(config Config) Logger {
	var output io.Writer = os.Stdout
	level := getZerologLevel(config.Level)

	zerolog.SetGlobalLevel(level)

	var logger zerolog.Logger
	if config.JSONFormat {
		logger = zerolog.New(output).With().
			Timestamp().
			Logger()
	} else {
		output = zerolog.ConsoleWriter{
			Out:        output,
			TimeFormat: "2006-01-02 15:04:05",
		}
		logger = zerolog.New(output).With().
			Timestamp().
			Logger()
	}

	return &consoleLogger{
		logger: &logger,
	}
}

func (l *consoleLogger) Debug(msg string, fields ...Field) {
	l.mu.Lock()
	e := l.logger.Debug()
	l.addFields(e, fields)
	e.Msg(msg)
	l.mu.Unlock()
}

func (l *consoleLogger) Info(msg string, fields ...Field) {
	l.mu.Lock()
	e := l.logger.Info()
	l.addFields(e, fields)
	e.Msg(msg)
	l.mu.Unlock()
}

func (l *consoleLogger) Warn(msg string, fields ...Field) {
	l.mu.Lock()
	e := l.logger.Warn()
	l.addFields(e, fields)
	e.Msg(msg)
	l.mu.Unlock()
}

func (l *consoleLogger) Error(msg string, fields ...Field) {
	l.mu.Lock()
	e := l.logger.Error()
	l.addFields(e, fields)
	e.Msg(msg)
	l.mu.Unlock()
}

func (l *consoleLogger) Fatal(msg string, fields ...Field) {
	l.mu.Lock()
	e := l.logger.Fatal()
	l.addFields(e, fields)
	e.Msg(msg)
	l.mu.Unlock()
}

func (l *consoleLogger) WithFields(fields ...Field) Logger {
	l.mu.Lock()
	e := l.logger.With()
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	newLogger := e.Logger()
	l.mu.Unlock()
	return &consoleLogger{
		logger: &newLogger,
		mu:     sync.Mutex{},
	}
}

func (l *consoleLogger) Sync() error {
	// No-op for console logger
	return nil
}

func (l *consoleLogger) addFields(e *zerolog.Event, fields []Field) {
	for _, f := range fields {
		e.Interface(f.Key, f.Value)
	}
}

func getZerologLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
