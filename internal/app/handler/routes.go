package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) RegisterAPI(r *fiber.App) {
	v1 := r.Group("/api")
	threats := v1.Group("/threats")
	threats.Get("/", h.ThreatAPI.GetThreats)
}
