package error_handler

import (
	"errors"
	"fmt"
	"github.com/compico/shopsite/internal/config"
	"github.com/compico/shopsite/pkg/errs"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"net/http"
)

type errorHandler struct {
	debugMode DebugMode
}

type DebugMode bool

func ProviderDebugMode(conf config.Config) DebugMode {
	return DebugMode(conf.IsDev())
}

func NewErrorHandler(debugMode DebugMode) fiber.ErrorHandler {
	errHandler := &errorHandler{debugMode: debugMode}
	return errHandler.handle
}

func (h errorHandler) handle(ctx fiber.Ctx, err error) error {
	code := http.StatusInternalServerError
	message := "Unknown error"
	var e errs.Error
	log.Errorf("%s", err)
	if errors.As(err, &e) {
		code = e.GetCode()
		message = e.Error()
	}

	if !h.debugMode {
		return ctx.Status(code).SendString(message)
	}

	return ctx.Status(code).SendString(fmt.Sprintf(
		"%s",
		err.Error(),
	))
}
