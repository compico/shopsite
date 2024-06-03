package config

import "time"

type (
	Http interface {
		GetAddr() string
		GetReadTimeout() time.Duration
		GetWriteTimeout() time.Duration
		GetStaticPath() string
	}

	HttpImpl struct {
		Addr         string        `yaml:"addr"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
		StaticPath   string        `yaml:"static_path"`
	}
)

func (hc *HttpImpl) GetAddr() string {
	return hc.Addr
}

func (hc *HttpImpl) GetReadTimeout() time.Duration {
	return hc.ReadTimeout
}

func (hc *HttpImpl) GetWriteTimeout() time.Duration {
	return hc.WriteTimeout
}

func (hc *HttpImpl) GetStaticPath() string {
	return hc.StaticPath
}
