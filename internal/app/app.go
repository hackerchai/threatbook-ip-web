package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
	"github.com/hackerchai/threatbook-ip-web/internal/app/global"
	"github.com/hackerchai/threatbook-ip-web/internal/app/handler"
	"github.com/hackerchai/threatbook-ip-web/internal/app/middleware"
	"time"
)

func New() *fiber.App {
	f := fiber.New()
	f.Use(middleware.AccessLogger(&middleware.AccessLoggerConfig{
		Logger:      global.Logger,
		Type:        global.CONFIG.Log.Type,
		Environment: global.CONFIG.Log.Environment,
		Filename:    global.CONFIG.Log.Filename,
		MaxSize:     global.CONFIG.Log.MaxSize,
		MaxAge:      global.CONFIG.Log.MaxAge,
		MaxBackups:  global.CONFIG.Log.MaxBackups,
		LocalTime:   global.CONFIG.Log.LocalTime,
		Compress:    global.CONFIG.Log.Compress,
		Level:       global.CONFIG.Log.GetLogLevel(),
	}))
	if global.CONFIG.CSRF.Enable {
		f.Use(csrf.New())
	}
	if global.CONFIG.Limiter.Enable {
		f.Use(limiter.New(limiter.Config{
			Max:          global.CONFIG.Limiter.Max,
			Expiration:   time.Duration(global.CONFIG.Limiter.Duration) * time.Minute,
			LimitReached: handler.OnLimitReached,
		}))
	}
	if global.CONFIG.CORS.Enable {
		f.Use(cors.New(cors.Config{
			AllowOrigins: global.CONFIG.CORS.AllowOrigins,
			AllowHeaders: global.CONFIG.CORS.AllowHeaders,
			AllowMethods: global.CONFIG.CORS.AllowMethods,
		}))
	}
	if global.CONFIG.Common.Swagger == "enable" {
		f.Get("/swagger/*", swagger.HandlerDefault)
	}
	f.Static("/", "./static/dist")
	return f
}
