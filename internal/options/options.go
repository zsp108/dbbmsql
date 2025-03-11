package options

import "time"

type Options struct {
	GeneralOptions   GeneralOptions   `yaml:"general"`
	MySQLOptions     MYSQLOptions     `yaml:"mysql"`
	OceanbaseOptions OceanbaseOptions `yaml:"oceanbase"`
}

type GeneralOptions struct {
	LogDir   string `yaml:"logdir" default:"dbbmsql.log"`
	LogLevel string `yaml:"loglevel" default:"info"`
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
