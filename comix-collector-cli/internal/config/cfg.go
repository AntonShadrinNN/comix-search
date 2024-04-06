package config

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Url    string `default:"https://xkcd.com/" yaml:"source_url"`
	DbFile string `default:"database.json" yaml:"db_file"`
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(c)
	type defaultCfg Config
	if err := unmarshal((*defaultCfg)(c)); err != nil {
		return err
	}

	return nil
}

func NewConfig() (Config, error) {
	yamlFile, err := os.ReadFile("config.yaml")
	if !os.IsNotExist(err) && err != nil {
		return Config{}, err
	}
	var c Config
	if len(yamlFile) == 0 {
		defaults.Set(&c)
		return c, nil
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
