package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

var instance Config

const (
	Dev = "dev"
)

type (
	Config interface {
		GetSiteName() string
		GetDescription() string
		GetTitle() string
		SetTitle(string)
		GetEnvironment() string
		GetTemplatePath() string
		GetTemplateExtension() string
		IsDev() bool
		GetHttpConfig() Http
		GetDatabaseConfig() Database
	}

	Impl struct {
		SiteName          string        `yaml:"site_name"`
		Description       string        `yaml:"description"`
		Title             string        `yaml:"title"`
		Environment       string        `yaml:"environment"`
		TemplatePath      string        `yaml:"template_path"`
		TemplateExtension string        `yaml:"template_extension"`
		Http              *HttpImpl     `yaml:"http_config"`
		Database          *DatabaseImpl `yaml:"database_config"`
	}
)

func NewConfig(path string) (Config, error) {
	var (
		err error
		f   []byte
	)

	if path == "" {
		return nil, fmt.Errorf("empty path")
	}

	if instance != nil {
		return instance, nil
	}

	sync.OnceFunc(func() {
		f, err = os.ReadFile(path)
		if err != nil {
			return
		}

		instance = &Impl{}

		err = yaml.Unmarshal(f, instance)
	})()

	return instance, err
}

func (c *Impl) GetSiteName() string {
	return c.SiteName
}

func (c *Impl) GetDescription() string {
	return c.Description
}

func (c *Impl) GetTitle() string {
	return c.Title
}

func (c *Impl) SetTitle(title string) {
	c.Title = title
}

func (c *Impl) GetEnvironment() string {
	return c.Environment
}

func (c *Impl) GetTemplatePath() string {
	return c.TemplatePath
}
func (c *Impl) GetTemplateExtension() string {
	return c.TemplateExtension
}

func (c *Impl) IsDev() bool {
	return c.Environment == Dev
}

func (c *Impl) GetHttpConfig() Http {
	return c.Http
}

func (c *Impl) GetDatabaseConfig() Database {
	return c.Database
}
