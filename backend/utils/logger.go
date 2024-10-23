package utils

import (
	"context"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

type CustomLogger struct {
	logger.Interface
	Log *log.Logger
}

func NewCustomLogger() *CustomLogger {
	var logLevel logger.LogLevel

	switch os.Getenv("LOG_LEVEL") {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Info
	}

	return &CustomLogger{
		Interface: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logLevel,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
		Log: log.New(os.Stdout, "\r\n", log.LstdFlags),
	}
}

func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	l.Log.Printf("[TRACE] SQL: %v\n", sql)
	l.Log.Printf("[TRACE] Rows affected: %d\n", rows)
	l.Log.Printf("[TRACE] Elapsed time: %v\n", elapsed)

	if err != nil {
		l.Log.Printf("[ERROR] %v\n", err)
	}
}

func (l *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.Interface = l.Interface.LogMode(level)
	return &newLogger
}

func (l *CustomLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.Log.Printf("[INFO] "+msg, data...)
}

func (l *CustomLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Log.Printf("[WARN] "+msg, data...)
}

func (l *CustomLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Log.Printf("[ERROR] "+msg, data...)
}
