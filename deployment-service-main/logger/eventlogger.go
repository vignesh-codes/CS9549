package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var EventLogger *zap.Logger

func InitEventLogger() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.EpochMillisTimeEncoder
	config.TimeKey = "timestamp"
	config.MessageKey = "event_source"
	config.LevelKey = ""
	fileEncoder := zapcore.NewJSONEncoder(config)

	// logger
	eventLogFile, _ := os.OpenFile("logs/events.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	eventLogLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel
	})
	eventLogWriter := zapcore.AddSync(eventLogFile)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, eventLogWriter, eventLogLevel),
	)

	EventLogger = zap.New(core)
	defer EventLogger.Sync()
}
