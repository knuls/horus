package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func (c *Config) SetFile(path string, name string, configType string) {
	c.v.AddConfigPath(path)
	c.v.SetConfigName(name)
	c.v.SetConfigType(configType)
}

func (c *Config) SetBindings(bindings []string) {
	for _, binding := range bindings {
		c.v.BindEnv(binding)
	}
}

func (c *Config) Load(o interface{}) error {
	if err := c.v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return err
		}
		return err
	}
	if err := c.v.Unmarshal(&o); err != nil {
		return err
	}
	c.v.AutomaticEnv()
	return nil
}

func New(prefix string) (*Config, error) {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return &Config{v: v}, nil
}
