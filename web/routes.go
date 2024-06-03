package web

import (
	"github.com/compico/shopsite/internal/handler"
	"github.com/compico/shopsite/internal/handler/page_render_handler"
	"github.com/gofiber/fiber/v3"
)

type RegisterRoutesResult bool

func RegisterRoutes(
	router fiber.Router,
	config ServerConfig,
	pageRenderHandler *page_render_handler.PageRender,
	index *handler.Index,
	getProducts *handler.GetProducts,
) RegisterRoutesResult {
	router.Static("/public", config.GetStaticPath())

	router.Get("/", pageRenderHandler.Handler)
	router.Get("/:page_name", pageRenderHandler.Handler)

	router.Get("/content/index", index.Handler)
	router.Get("/products/:page", getProducts.Handler)

	return true
}
