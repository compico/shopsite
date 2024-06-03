package web

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"time"
)

type ServerImpl struct {
	App  *fiber.App
	Conf ServerConfig
}

type Server interface {
	Router() fiber.Router
	Start()
	Stop()
}

type ServerConfig interface {
	GetAddr() string
	GetReadTimeout() time.Duration
	GetWriteTimeout() time.Duration
	GetStaticPath() string
}

func NewServer(conf ServerConfig, errorHandler fiber.ErrorHandler, render fiber.Views) Server {
	return &ServerImpl{
		App: fiber.New(fiber.Config{
			ReadTimeout:  conf.GetReadTimeout(),
			WriteTimeout: conf.GetWriteTimeout(),
			ErrorHandler: errorHandler,
			Views:        render,
		}),
		Conf: conf,
	}
}

func (s *ServerImpl) Start() {
	s.App.Use(func(ctx fiber.Ctx) error {
		start := time.Now().UTC()
		defer func() {
			duration := time.Now().UTC().Sub(start)
			if duration > 1*time.Second {
				log.WithContext(ctx.Context()).Warnf(
					"Long request: %+v",
					struct {
						Path     string
						Method   string
						Duration time.Duration
					}{
						Path:     string(ctx.Request().URI().Path()),
						Method:   ctx.Method(),
						Duration: duration,
					},
				)
			}
		}()
		return ctx.Next()
	})

	if err := s.App.Listen(s.Conf.GetAddr()); err != nil {
		log.DefaultLogger().Errorf("error on shutdown")
	}
}

func (s *ServerImpl) Stop() {
	if err := s.App.Shutdown(); err != nil {
		log.DefaultLogger().Errorf("error on shutdown")
	}
}

func (s *ServerImpl) Router() fiber.Router {
	return s.App
}
