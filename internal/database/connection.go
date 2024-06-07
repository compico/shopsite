package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var instance *gorm.DB
var once *sync.Once

type (
	ConnectionResult bool
	Config           interface {
		GetDSN() string
		AddParams(...string)
		GetDialect() string
	}
)

func NewConnection(conf Config) (*gorm.DB, error) {
	if instance != nil {
		return instance, nil
	}

	var err error

	once.Do(func() {
		instance, err = gorm.Open(
			mysql.Open(conf.GetDSN()),
			&gorm.Config{
				Logger: logger.New(log.New(os.Stdout, "[DB]", log.LstdFlags), logger.Config{
					SlowThreshold:             10 * time.Second,
					Colorful:                  false,
					IgnoreRecordNotFoundError: true,
					ParameterizedQueries:      true,
					LogLevel:                  logger.Info,
				}),
			},
		)
	})

	return instance, err
}

func GetInstance() *gorm.DB {
	return instance
}
