package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// 自定义日志结构体
type Logger struct {
	*logrus.Logger
}

/*
NewLogger 创建一个新的 Logger 实例并设置格式、级别、输出位置
级别：DEBUG，INFO，WARN，ERROR
格式：json，text
输出位置：stdout，stderr，file
*/
func NewLogger(level logrus.Level, format string, output string) *Logger {
	logger := logrus.New()

	// 设置日志级别
	logger.SetLevel(level)

	// 设置日志输出位置
	switch output {
	case "stdout":
		logger.SetOutput(os.Stdout)
	case "stderr":
		logger.SetOutput(os.Stderr)
	default:
		// 文件输出
		file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.SetOutput(file)
		} else {
			logger.SetOutput(os.Stdout) // 如果文件打开失败则输出到标准输出
		}
	}

	// 设置日志格式
	switch format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	default:
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return &Logger{logger}
}

// 日志级别快捷方法
func (l *Logger) Debug(v ...interface{}) {
	l.Logger.Debug(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.Logger.Info(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.Logger.Warn(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.Logger.Error(v...)
}
