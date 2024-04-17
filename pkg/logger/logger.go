package logger

import (
	"go.uber.org/zap"
	"log"
)

func SetupLogger() (*zap.Logger, error) {
	zapLogger, err := zap.NewDevelopment()

	if err != nil {
		log.Fatal("cant initialize logger %w", err)
	}

	defer zapLogger.Sync()

	zap.ReplaceGlobals(zapLogger)

	return zapLogger, nil
}
