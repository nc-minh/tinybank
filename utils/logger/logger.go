package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func newLogger() (*zap.Logger, error) {
	encoderConfig := setDebugMode(true)
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	return zap.New(core), nil
}

func setDebugMode(isDebug bool) zapcore.EncoderConfig {
	if isDebug {
		return zap.NewDevelopmentConfig().EncoderConfig
	} else {
		return zap.NewProductionConfig().EncoderConfig
	}
}

// Info logs a message at level Info on the standard logger.
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Error logs a message at level Error on the standard logger.
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)

}

// Debug logs a message at level Debug on the standard logger.
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)

}

// Warn logs a message at level Warn on the standard logger.
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)

}

// init initializes the logger
func init() {
	logger, _ = newLogger()
}
