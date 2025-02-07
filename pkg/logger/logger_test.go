package logger_test

import (
	"testing"

	"github.com/zsp108/dbbmsql/pkg/logger"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	log := logger.NewLogger(logrus.DebugLevel, "text", "stdout")

	// 使用日志
	log.Debug("This is a debug message")
	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")
}
