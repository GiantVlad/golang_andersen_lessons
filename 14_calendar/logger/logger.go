package logger

import (
	"go.uber.org/zap"
	"log"
)

func NewLogger(logFile string, level string) *zap.SugaredLogger {
	zcfg := zap.NewProductionConfig()
	if logFile == "" {
		logFile = "stdout"
	} else {
		logFile = "./" + logFile
	}
	zcfg.OutputPaths = []string{
		logFile,
	}
	var err error
	zcfg.Level, err = zap.ParseAtomicLevel(level)
	if err != nil {
		log.Fatalf("Invalid log level: %s", level)
	}
	logger, err := zcfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	return logger.Sugar()
}
