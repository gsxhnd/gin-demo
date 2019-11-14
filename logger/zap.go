package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger01 = logger()

func logger() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		os.Stdout,
		zap.DebugLevel,
	)
	logger := zap.New(core)
	defer logger.Sync() // flushes buffer, if any
	logger.Info("DefaultLogger init success")
	return logger
}
