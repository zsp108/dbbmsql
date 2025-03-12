package options

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zsp108/dbbmsql/pkg/logger"
)

type Options struct {
	GeneralOptions   GeneralOptions   `yaml:"general"`
	MySQLOptions     MYSQLOptions     `yaml:"mysql"`
	OceanbaseOptions OceanbaseOptions `yaml:"oceanbase"`
}

type GeneralOptions struct {
	LogOutPut string `yaml:"logoutput" default:"dbbmsql.log"`
	LogLevel  string `yaml:"loglevel" default:"info"`
	LogType   string `yaml:"logtype" default:"text"`
}

type MYSQLOptions struct {
	Host                  string        `yaml:"host"`
	Username              string        `yaml:"username"`
	Password              string        `yaml:"password"`
	Port                  int           `yaml:"port"`
	Database              string        `yaml:"database"`
	MaxIdleConnections    int           `yaml:"max_idle_connections" default:"10"`
	MaxOpenConnections    int           `yaml:"max_open_connections" default:"100"`
	MaxConnectionLifeTime time.Duration `yaml:"max_connection_life_time" default:"10s"`
}

type OceanbaseOptions struct {
	Host                  string        `yaml:"host"`
	Username              string        `yaml:"username"`
	Password              string        `yaml:"password"`
	Port                  int           `yaml:"port"`
	Database              string        `yaml:"database"`
	MaxIdleConnections    int           `yaml:"max_idle_connections"`
	MaxOpenConnections    int           `yaml:"max_open_connections"`
	MaxConnectionLifeTime time.Duration `yaml:"max_connection_life_time"`
}

// 构造函数，设置日志默认值
func NewLoggerOptions() *GeneralOptions {
	return &GeneralOptions{
		LogOutPut: "stdout", // 默认值
		LogLevel:  "info",   // 默认值
		LogType:   "text",   // 默认值
	}
}

// 构造函数，设置Mysql默认值
func NewMYSQLOptions() *MYSQLOptions {
	return &MYSQLOptions{
		MaxIdleConnections:    10,            // 默认值
		MaxOpenConnections:    100,           // 默认值
		MaxConnectionLifeTime: 1 * time.Hour, // 默认值
	}
}

// 构造函数，设置OB默认值
func NewOceanbaseOptions() *MYSQLOptions {
	return &MYSQLOptions{
		MaxIdleConnections:    10,            // 默认值
		MaxOpenConnections:    100,           // 默认值
		MaxConnectionLifeTime: 1 * time.Hour, // 默认值
	}
}

func NewLogger(opts *Options) *logger.Logger {

	level, err := logrus.ParseLevel(opts.GeneralOptions.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger := logger.NewLogger(level, opts.GeneralOptions.LogType, opts.GeneralOptions.LogOutPut)

	return logger
}
