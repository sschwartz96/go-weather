package main

import (
	"fmt"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

type config struct {
	ApiKey string `yaml:"apiKey"`
	CityID string `yaml:"cityID"`
	Units  string `yaml:"units"`
}

func loadConfig() (*config, error) {
	// open user for home directory
	user, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("LoadConfig(), cannot find user home: %v", err)
	}

	// open config file
	configLoc := fmt.Sprintf("%s/.config/go-weather/config.yml", user.HomeDir)
	file, err := os.Open(configLoc)
	if err != nil {
		return nil, fmt.Errorf("LoadConfig(), cannot open config file: %v", err)
	}

	// decode file contents
	cfg := &config{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, fmt.Errorf("LoadConfig(), cannot decode config file: %v", err)
	}

	return cfg, nil
}
