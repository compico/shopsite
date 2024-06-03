package di

import (
	"github.com/compico/shopsite/internal/config"
	"github.com/compico/shopsite/internal/database"
	"github.com/compico/shopsite/web"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
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

func ProviderConnectToDatabase(conf database.Config) (database.ConnectionResult, error) {
	_, err := database.NewConnection(conf)
	if err != nil {
		return false, err
	}
	return true, nil
}
