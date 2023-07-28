package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Output  *zap.Logger
	Console *zap.Logger
)

func InitializeLoggers(logLevel string) {
	newOutputLogger(logLevel)
	newConsoleLogger(logLevel)
}

func newOutputLogger(logLevel string) {
	logConfig := zap.Config{
		Level:            zap.NewAtomicLevel(),
		Encoding:         "json",
		OutputPaths:      []string{"stdout", "./output.log"},
		ErrorOutputPaths: []string{"stderr", "./output.log"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			LevelKey:     "level",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		fmt.Println("[INFO] Using default DEBUG log level for OUTPUT logger")
		level = zap.DebugLevel
	}

	logConfig.Level.SetLevel(level)

	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	Output = logger

	defer Output.Sync()
}

func newConsoleLogger(logLevel string) {
	logConfig := zap.Config{
		Level:            zap.NewAtomicLevel(),
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			LevelKey:     "level",
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		fmt.Println("[INFO] Using default DEBUG log level for CONSOLE logger")
		level = zap.DebugLevel
	}

	logConfig.Level.SetLevel(level)

	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	Console = logger

	defer Console.Sync()
}
