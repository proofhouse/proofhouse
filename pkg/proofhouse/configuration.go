package proofhouse

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config structure holds all configuration parameters.
type Config struct {
	FeaturesDir string   `yaml:"features_dir"`
	Plugins     []string `yaml:"plugins"`
}

// NewConfig creates new Config struct with values parsed from "path" configuration file.
func NewConfig(path string) (c *Config, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to read configuration file '%v'", path))
	}

	c = new(Config)

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, errors.Wrap(err, "Error occurred while parsing yaml configuration file")
	}

	if err := c.validate(); err != nil {
		return nil, errors.Wrap(err, "Configuration validation failed")
	}

	return c, nil
}

// validate mapped configuration parameters.
func (c *Config) validate() (err error) {
	stat, err := os.Stat(c.FeaturesDir)
	if err != nil {
		return errors.Wrapf(err, "Failed to get statistics of the directory containing features '%v'", err)
	}
	if stat.IsDir() == false {
		return errors.Errorf("Features path '%s' is not a directory", c.FeaturesDir)
	}

	return nil
}
