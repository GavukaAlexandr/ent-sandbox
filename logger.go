package main

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger    *zap.Logger
	zapConfig zap.Config
)

func init() {
	logLevel := os.Getenv("LOG_LEVEL")
	if len(logLevel) == 0 {
		zap.L().Fatal("log conf only allow [debug, info, warn, error], please check your configure.")
	}
	logLevel = strings.ToLower(logLevel)

	setLogLevel(logLevel)
	if logLevel == "debug" {
		zapConfig.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		zapConfig.EncoderConfig = zap.NewProductionEncoderConfig()
	}

	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.OutputPaths = []string{"stderr"}
	zapConfig.ErrorOutputPaths = []string{"stderr"}
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConfig.Encoding = "console"

	Logger, _ = zapConfig.Build()
	zap.ReplaceGlobals(Logger)
}

func setLogLevel(level string) {
	switch level {
	case "debug":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		break
	case "info":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		break
	case "warn":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
		break
	case "error":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
		break
	default:
		zap.L().Fatal("log conf only allow [debug, info, warn, error], please check your configure.")
	}
}
