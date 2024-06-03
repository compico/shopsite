package config

import (
	"fmt"
	"strings"
)

type (
	Database interface {
		GetDSN() string
		GetDialect() string
		GetUser() string
		GetPassword() string
		GetAddress() string
		GetPort() uint16
		GetDBName() string
		GetParams() []string
		AddParams(...string)
	}

	DatabaseImpl struct {
		Dialect  string   `yaml:"dialect"`
		User     string   `yaml:"user"`
		Password string   `yaml:"password"`
		Address  string   `yaml:"address"`
		Port     uint16   `yaml:"port"`
		DBName   string   `yaml:"db_name"`
		Params   []string `yaml:"params"`
	}
)

func (c *DatabaseImpl) GetDSN() string {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		c.User,
		c.Password,
		c.Address,
		c.Port,
		c.DBName,
		strings.Join(c.Params, "&"),
	)
}

func (c *DatabaseImpl) GetDialect() string {
	return c.Dialect
}

func (c *DatabaseImpl) GetUser() string {
	return c.User
}

func (c *DatabaseImpl) GetPassword() string {
	return c.Password
}

func (c *DatabaseImpl) GetAddress() string {
	return c.Address
}

func (c *DatabaseImpl) GetPort() uint16 {
	return c.Port
}

func (c *DatabaseImpl) GetDBName() string {
	return c.DBName
}

func (c *DatabaseImpl) GetParams() []string {
	return c.Params
}

func (c *DatabaseImpl) AddParams(params ...string) {
	c.Params = append(c.Params, params...)
}
