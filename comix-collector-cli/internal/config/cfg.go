package config

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

// Config is a configuration struct for app
type Config struct {
	Url    string `default:"https://xkcd.com/" yaml:"source_url"` // url to fetch comixes from
	DbFile string `default:"database.json" yaml:"db_file"`        // database file
}

// UnmarshalYAML overrides default behaviour of yaml unmarshaller and makes it possible
// to set default values on fields
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := defaults.Set(c)
	if err != nil {
		return err
	}
	type defaultCfg Config
	if err := unmarshal((*defaultCfg)(c)); err != nil {
		return err
	}

	return nil
}

// NewConfig parses configuration file and returns Config object
func NewConfig() (Config, error) {
	yamlFile, err := os.ReadFile("config.yaml")
	if !os.IsNotExist(err) && err != nil {
		return Config{}, err
	}
	var c Config
	if len(yamlFile) == 0 {
		err = defaults.Set(&c)
		if err != nil {
			return Config{}, err
		}
		return c, nil
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
