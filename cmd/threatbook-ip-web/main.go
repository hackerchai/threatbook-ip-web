package main

import (
	"fmt"
	"github.com/hackerchai/threatbook-ip-web/internal/app"
	"github.com/hackerchai/threatbook-ip-web/internal/app/global"
	_ "github.com/hackerchai/threatbook-ip-web/internal/app/swagger"
	"github.com/hackerchai/threatbook-ip-web/pkg/logger"
)

// @title threatbook-ip-web
// @version 0.1.0
// @description Web for craling threatbook community to obtain threat ip from feeds.
// @schemes http https
// @basePath /
// @contact.name Eason Chai
// @contact.email i@hackerchai.com
func main() {
	logger.Init()
	logger.Info("threatbook-ip-web start")
	app.InitConfig()
	app.InitDatabase()
	injector := app.InitializeEvent()
	injector.Handler.RegisterAPI(injector.Engine)
	address := fmt.Sprintf("%s:%d", global.CONFIG.Common.Host, global.CONFIG.Common.Port)
	err := injector.Engine.Listen(address)
	if err != nil {
		logger.Panic("Start Fiber Server error: " + err.Error())
	}
}
