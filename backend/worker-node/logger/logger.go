package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger creates and configures a new zap logger
func InitLogger(logLevel string) (*zap.Logger, error) {
	// Parse log level
	var level zapcore.Level
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		// Default to info if parsing fails
		level = zapcore.InfoLevel
	}
	
	// Create encoder configuration with colors and formatting
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,  // Colors for different log levels
		EncodeTime:     zapcore.ISO8601TimeEncoder,        // ISO8601 time format
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	
	// Create console encoder with our custom config
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	
	// Create core with stdout as output
	core := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		level,
	)
	
	// Create logger with our core
	logger := zap.New(core)
	
	return logger, nil
}