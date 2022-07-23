package config

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Common       Common   `mapstructure:"common"`
	DevDatabase  Database `mapstructure:"dev_database"`
	ProdDatabase Database `mapstructure:"prod_database"`
	Log          Log      `mapstructure:"log"`
	CORS         CORS     `mapstructure:"cors"`
	Limiter      Limiter  `mapstructure:"limiter"`
	CSRF         CSRF     `mapstructure:"csrf"`
}

type Common struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	DeployMode string `mapstructure:"deploy_mode"`
	Swagger    string `mapstructure:"swagger"`
}

type Database struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type Log struct {
	Type       string `mapstructure:"type"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	LocalTime  bool   `mapstructure:"local_time"`
	Compress   bool   `mapstructure:"compress"`
	Level      string `mapstructure:"level"`
}

type CORS struct {
	Enable       bool   `mapstructure:"enable"`
	AllowOrigins string `mapstructure:"allow_origins"`
	AllowHeaders string `mapstructure:"allow_headers"`
	AllowMethods string `mapstructure:"allow_methods"`
}

type Limiter struct {
	Enable   bool `mapstructure:"enable"`
	Max      int  `mapstructure:"limit_max"`
	Duration int  `mapstructure:"limit_duration"`
}

type CSRF struct {
	Enable bool `mapstructure:"enable"`
}

func (d *Database) Dsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", d.Host, d.Port, d.Username, d.Database, d.Password, d.SSLMode)
}

func (l *Log) GetLogLevel() zapcore.Level {
	switch l.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
