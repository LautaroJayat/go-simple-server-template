package config

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ReadTimeout    int    `yaml:"ReadTimeout"`
	WriteTimeout   int    `yaml:"WriteTimeout"`
	MaxHeaderBytes int    `yaml:"MaxHeaderBytes"`
	Port           string `yaml:"Port"`
}

func FromYAML(f io.Reader) (*Config, error) {
	b, err := io.ReadAll(f)

	if err != nil {
		return nil, fmt.Errorf("couldn't read config file. The error was %q", err)
	}
	cfg := &Config{}
	err = yaml.Unmarshal(b, cfg)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshal yaml file: %q", err)
	}
	return cfg, nil

}
