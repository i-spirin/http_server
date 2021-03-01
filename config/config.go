package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents all configuration of the application
type Config struct {
	BindHost string `yaml:"bind_host"`
	BindPort int    `yaml:"bind_port"`
}

// Parse is function to get configuration from YAML-file
func (c *Config) Parse(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	err = yaml.NewDecoder(file).Decode(c)
	if err != nil {
		return err
	}
	return nil
}
