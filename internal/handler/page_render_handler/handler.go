package page_render_handler

import (
	"github.com/compico/shopsite/internal/config"
	"github.com/gofiber/fiber/v3"
)

type PageRender struct {
	Config config.Config
}

func NewPageRender(conf config.Config) *PageRender {
	return &PageRender{
		Config: conf,
	}
}

func (h *PageRender) Handler(ctx fiber.Ctx) error {
	var title string
	pageName := ctx.Params("page_name", "index")

	switch pageName {
	case "index":
		title = "Главная страница"
	case "products":
		title = "Товары"
	default:
		title = ""
	}

	if title != "" {
		title = title + " - " + h.Config.GetTitle()
	}

	if title == "" {
		title = h.Config.GetTitle()
	}

	return ctx.Render(
		"common/document",
		map[string]any{
			"Title":    title,
			"PageName": pageName,
		},
	)
}
