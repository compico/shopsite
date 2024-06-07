//go:build wireinject
// +build wireinject

package migrate

import (
	"database/sql"
	"github.com/compico/shopsite/internal/database"
	"github.com/compico/shopsite/internal/di"
	"github.com/google/wire"
)

type (
	migrator struct {
		conf       Config
		connection *sql.DB
	}

	Config interface {
		GetDSN() string
		AddParams(...string)
		GetDialect() string
	}
)

func InitializeMigrator(configPath string, dsnArgs di.DsnParams) (migrator, error) {
	panic(wire.Build(
		di.ProviderConfig,
		di.ProviderDataBaseConfigWithParams,

		database.NewConnection,
		di.ProviderDatabaseTx,

		wire.Bind(new(Config), new(database.Config)),

		wire.Struct(new(migrator), "*")))
}
