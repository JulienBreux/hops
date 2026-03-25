package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Check struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Endpoint string `yaml:"endpoint"`
	Method   string `yaml:"method"`
}

type Service struct {
	Name        string  `yaml:"name"`
	Description string  `yaml:"description"`
	Emoji       string  `yaml:"emoji"`
	Checks      []Check `yaml:"checks"`
}

type Config struct {
	Services []Service `yaml:"services"`
}

func Parse(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	expanded := os.ExpandEnv(string(data))

	var config Config
	if err := yaml.Unmarshal([]byte(expanded), &config); err != nil {
		return nil, err
	}

	return &config, nil
}
