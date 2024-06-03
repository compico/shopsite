package handler

import "github.com/gofiber/fiber/v3"

type GetProducts struct {
}

func NewGetProducts() *GetProducts {
	return &GetProducts{}
}

func (h *GetProducts) Handler(ctx fiber.Ctx) error {
	return ctx.JSON(fiber.Map{})
}
