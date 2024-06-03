package handler

import (
	"github.com/compico/shopsite/internal/config"
	"github.com/gofiber/fiber/v3"
)

type Index struct {
	Config config.Config
}

func NewIndex(conf config.Config) *Index {
	return &Index{
		Config: conf,
	}
}

func (h *Index) Handler(ctx fiber.Ctx) error {
	return ctx.Render("index", map[string]any{
		"Description": h.Config.GetDescription(),
	})
}
