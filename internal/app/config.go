package app

import (
	"github.com/hackerchai/threatbook-ip-web/internal/app/global"
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
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Read config failed: ", err)
	}
	if err := viper.Unmarshal(&global.CONFIG); err != nil {
		log.Fatal("Unmarshal config failed: ", err)
	}
}
