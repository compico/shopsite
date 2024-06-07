//go:build wireinject
// +build wireinject

package http

import (
	"github.com/compico/shopsite/internal/database"
	"github.com/compico/shopsite/internal/di"
	"github.com/compico/shopsite/internal/handler"
	"github.com/compico/shopsite/internal/handler/error_handler"
	"github.com/compico/shopsite/internal/handler/page_render_handler"
	"github.com/compico/shopsite/web"
	"github.com/google/wire"
)

type (
	App struct {
		server web.Server
		rr     web.RegisterRoutesResult
		dr     database.ConnectionResult
	}
)

func InitializeApp(configPath string) (*App, error) {
	panic(wire.Build(
		di.ProviderConfig,
		di.ProviderServerConfig,
		di.ProviderDatabaseConfig,
		di.ProviderHtmlRender,
		di.ProviderDatabaseConnection,

		error_handler.ProviderDebugMode,
		error_handler.NewErrorHandler,

		page_render_handler.NewPageRender,
		handler.NewIndex,
		handler.NewGetProducts,

		web.NewServer,
		web.RegisterRoutes,
		di.ProviderRouter,
		wire.Struct(new(App), "*"),
	))
}
