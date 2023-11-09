package logger

import (
	"github.com/test_service/pkg/logger"
	"go.uber.org/zap"
)

// Log is a package level variable, every program should access logging function through "Log"
var (
	Log *zap.Logger
)

// SetLogger is the setter for log variable, it should be the only way to assign value to log
func SetLogger(cfg *logger.LoggingConfig) {
	Log = logger.New(cfg)
}
