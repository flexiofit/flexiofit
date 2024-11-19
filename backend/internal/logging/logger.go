// internal/logging/logger.go
// internal/logging/logger.go
package logging

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerConfig struct {
	Level       string
	Filepath    string
	MaxSize     int
	MaxBackups  int
	MaxAge      int
	Compress    bool
	Development bool
}

// CreateCustomEncoder creates a more detailed and readable encoder
func createCustomEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncodeTime = customTimeEncoder
	config.EncodeDuration = zapcore.StringDurationEncoder
	config.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewConsoleEncoder(config)
}

// customTimeEncoder formats timestamp with millisecond precision
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000 MST"))
}

// NewLogger creates a highly configurable zap logger
func NewLogger(config LoggerConfig) (*zap.Logger, error) {
	// Determine log level
	var level zapcore.Level
	switch config.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	// Encoder configuration
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
	}

	// Create log rotation
	var logWriter zapcore.WriteSyncer
	if config.Filepath != "" {
		logRotator := &lumberjack.Logger{
			Filename:   config.Filepath,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		}
		logWriter = zapcore.AddSync(logRotator)
	} else {
		logWriter = zapcore.AddSync(os.Stdout)
	}

	// Create core with multiple outputs
	core := zapcore.NewCore(
		createCustomEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(logWriter, zapcore.AddSync(os.Stdout)),
		zap.NewAtomicLevelAt(level),
	)

	// Logger options
	opts := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	}

	// Development mode
	if config.Development {
		opts = append(opts, zap.Development())
	}

	return zap.New(core, opts...), nil
}

// Global logger wrapper with enhanced functionality
var (
	Log   *zap.Logger
	Sugar *zap.SugaredLogger
)

// InitializeLogger sets up the global logger
func InitializeLogger(config LoggerConfig) error {
	var err error
	Log, err = NewLogger(config)
	if err != nil {
		return err
	}
	Sugar = Log.Sugar()
	return nil
}

// Sync flushes any buffered log entries
func Sync() error {
	if err := Log.Sync(); err != nil {
		return err
	}
	return Sugar.Sync()
}

// WithFields creates a contextual logger with additional fields
func WithFields(fields ...zap.Field) *zap.Logger {
	return Log.With(fields...)
}
