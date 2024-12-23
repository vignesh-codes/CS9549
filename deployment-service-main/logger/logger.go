package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var ConsoleLogger *zap.Logger
var LOG_CHECK bool = true

func InitLogger() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.LevelKey = "level"
	config.MessageKey = "event"
	config.StacktraceKey = ""
	config.TimeKey = "time"
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile("logs/api_info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.DebugLevel
	})
	// error and fatal level enabler
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.ErrorLevel
	})
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	errLogFile, _ := os.OpenFile("logs/api_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	errWriter := zapcore.AddSync(errLogFile)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, infoLevel),
		zapcore.NewCore(fileEncoder, errWriter, errorLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	Logger = zap.New(core)
	Logger.Info("Initialized Logger")
	ConsoleLogger, _ = zap.NewProduction()
	defer Logger.Sync()
	defer ConsoleLogger.Sync()
}

func InitConsoleLogger() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	// Configure zap's encoder with ISO8601 time format
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.LevelKey = "level"
	config.MessageKey = "event"
	config.StacktraceKey = ""
	config.TimeKey = "time"

	// Create a console encoder that outputs logs to the console
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	// Create a core that will log to the console only
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	// Initialize the logger with the core
	Logger = zap.New(core)

	// Log an info message to verify logger initialization
	Logger.Info("Logger initialized for console logging")
	ConsoleLogger, _ = zap.NewProduction()
	// Ensure that any buffered logs are written to the console
	defer Logger.Sync()
	defer ConsoleLogger.Sync()
}
