// internal/logging/logger.go
package logging

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[37m"
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

// customTimeEncoder formats timestamp with millisecond precision and green color
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	timeStr := fmt.Sprintf("%s%s%s",
		colorGreen, // Changed from colorGray to colorGreen
		t.Format("2006-01-02 15:04:05.000 MST"),
		colorReset,
	)
	enc.AppendString(timeStr)
}

// customLevelEncoder adds colors to log levels
func customLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var levelColor string
	switch l {
	case zapcore.DebugLevel:
		levelColor = colorPurple
	case zapcore.InfoLevel:
		levelColor = colorBlue
	case zapcore.WarnLevel:
		levelColor = colorYellow
	case zapcore.ErrorLevel:
		levelColor = colorRed
	default:
		levelColor = colorReset
	}
	enc.AppendString(fmt.Sprintf("%s%s%s", levelColor, l.String(), colorReset))
}

// customCallerEncoder formats the caller path with color
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%s%s%s", colorCyan, caller.TrimmedPath(), colorReset))
}

// createCustomEncoder creates a more detailed and readable encoder with colors
func createCustomEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	config.EncodeLevel = customLevelEncoder
	config.EncodeTime = customTimeEncoder
	config.EncodeCaller = customCallerEncoder
	config.EncodeDuration = zapcore.StringDurationEncoder
	config.MessageKey = "message"
	config.LevelKey = "level"
	config.TimeKey = "timestamp"
	config.CallerKey = "caller"
	config.NameKey = "logger"
	config.StacktraceKey = "stacktrace"
	config.LineEnding = zapcore.DefaultLineEnding

	return zapcore.NewConsoleEncoder(config)
}

// NewLogger creates a highly configurable zap logger
func NewLogger(config LoggerConfig) (*zap.Logger, error) {
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

	core := zapcore.NewCore(
		createCustomEncoder(zapcore.EncoderConfig{}),
		zapcore.NewMultiWriteSyncer(logWriter, zapcore.AddSync(os.Stdout)),
		zap.NewAtomicLevelAt(level),
	)

	opts := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	}

	if config.Development {
		opts = append(opts, zap.Development())
	}

	return zap.New(core, opts...), nil
}

var (
	Log   *zap.Logger
	Sugar *zap.SugaredLogger
)

func InitializeLogger(config LoggerConfig) error {
	var err error
	Log, err = NewLogger(config)
	if err != nil {
		return err
	}
	Sugar = Log.Sugar()
	return nil
}

func Sync() error {
	if err := Log.Sync(); err != nil {
		return err
	}
	return Sugar.Sync()
}

func WithFields(fields ...zap.Field) *zap.Logger {
	return Log.With(fields...)
}
