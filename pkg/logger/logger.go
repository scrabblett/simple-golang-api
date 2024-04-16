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

	defer func(zapLogger *zap.Logger) {
		err = zapLogger.Sync()
		if err != nil {
			log.Fatal("cant flush logger %w", err)
		}
	}(zapLogger)

	zap.ReplaceGlobals(zapLogger)

	return zapLogger, nil
}
