package app

import (
	"github.com/hackerchai/threatbook-ip-web/internal/app/global"
	"github.com/hackerchai/threatbook-ip-web/pkg/logger"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	viper.SetDefault("common.deploy_mode", "development")
	viper.SetDefault("common.host", "127.0.0.1")
	viper.SetDefault("common.port", "8080")
	viper.SetDefault("common.swagger", "disable")
	viper.SetDefault("log.type", "console")
	viper.SetDefault("log.filename", "./log/app.log")
	viper.SetDefault("log.max_size", "100")
	viper.SetDefault("log.max_age", "30")
	viper.SetDefault("log.max_backups", "10")
	viper.SetDefault("log.local_time", "true")
	viper.SetDefault("log.compress", "true")
	viper.SetDefault("log.level", "info")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("Read config failed: " + err.Error())
	}
	if err := viper.Unmarshal(&global.CONFIG); err != nil {
		log.Fatal("Unmarshal config failed: " + err.Error())
	}
}
