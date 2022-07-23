package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hackerchai/threatbook-ip-web/internal/app/web"
	"github.com/hackerchai/threatbook-ip-web/pkg/errors"
)

func (h *Handler) RegisterAPI(r *fiber.App) {
	v1 := r.Group("/api")
	threats := v1.Group("/threats")
	threats.Get("/", h.ThreatAPI.GetThreats)
}

func OnLimitReached(c *fiber.Ctx) error {
	return web.ResError(c, errors.ErrTooManyRequests)
}
