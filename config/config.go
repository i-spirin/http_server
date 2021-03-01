package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents all configuration of the application
type Config struct {
	BindHost string `yaml:"bind_host"`
	BindPort int    `yaml:"bind_port"`
}

// Parse is function to get configuration from YAML-file
func (c *Config) Parse(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open file: %s", filename)
	}

	defer file.Close()

	err = yaml.NewDecoder(file).Decode(c)
	if err != nil {
		log.Fatalf("Could not parse YAML-file: %v", err)
	}
}
