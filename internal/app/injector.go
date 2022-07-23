package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/hackerchai/threatbook-ip-web/internal/app/handler"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine  *fiber.App
	Handler *handler.Handler
}
