package di

import (
	"database/sql"
	"github.com/compico/shopsite/internal/config"
	"github.com/compico/shopsite/internal/database"
	"github.com/compico/shopsite/web"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"gorm.io/gorm"
)

type (
	DsnParams []string
)

func ProviderConfig(configPath string) (config.Config, error) {
	return config.NewConfig(configPath)
}

func ProviderServerConfig(conf config.Config) web.ServerConfig {
	return conf.GetHttpConfig()
}

func ProviderHtmlRender(conf config.Config) fiber.Views {
	return html.New(conf.GetTemplatePath(), conf.GetTemplateExtension())
}

func ProviderRouter(server web.Server) fiber.Router {
	return server.Router()
}

func ProviderDatabaseConfig(conf config.Config) database.Config {
	return conf.GetDatabaseConfig()
}

func ProviderDataBaseConfigWithParams(conf config.Config, params DsnParams) database.Config {
	c := conf.GetDatabaseConfig()
	c.AddParams(params...)

	return c
}

func ProviderDatabaseConnection(conf database.Config) (database.ConnectionResult, error) {
	_, err := database.NewConnection(conf)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ProviderDatabaseTx(conn *gorm.DB) (*sql.DB, error) {
	return conn.DB()
}
