package config

func InitConfig(title string, description string) *Config {
	x := new(Config)
	x.Title = title
	x.Description = description
	return x
}

func (c *Config) SetTitle(t string) {
	c.Title = t
}
