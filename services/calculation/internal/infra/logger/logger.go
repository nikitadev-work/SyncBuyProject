package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

type LoggerInterface interface {
	Debug(msg string, params map[string]interface{})
	Info(msg string, params map[string]interface{})
	Warn(msg string, params map[string]interface{})
	Error(msg string, params map[string]interface{})
}

type Logger struct {
	logger *slog.Logger
}

var _ LoggerInterface = (*Logger)(nil)

func NewLogger(levelStr string) *Logger {
	var level slog.Level

	switch strings.ToUpper(levelStr) {
	case "DEBUG":
		level = slog.LevelDebug
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))

	return &Logger{
		logger: l,
	}
}

func (l *Logger) Debug(msg string, params map[string]interface{}) {
	l.logger.LogAttrs(context.Background(), slog.LevelDebug, msg, paramsToAttr(params)...)
}

func (l *Logger) Info(msg string, params map[string]interface{}) {
	l.logger.LogAttrs(context.Background(), slog.LevelInfo, msg, paramsToAttr(params)...)
}

func (l *Logger) Warn(msg string, params map[string]interface{}) {
	l.logger.LogAttrs(context.Background(), slog.LevelWarn, msg, paramsToAttr(params)...)
}

func (l *Logger) Error(msg string, params map[string]interface{}) {
	l.logger.LogAttrs(context.Background(), slog.LevelError, msg, paramsToAttr(params)...)
}

func paramsToAttr(params map[string]interface{}) []slog.Attr {
	if len(params) == 0 {
		return nil
	}

	attr := make([]slog.Attr, 0, len(params))

	for k, v := range params {
		attr = append(attr, slog.Any(k, v))
	}

	return attr
}
