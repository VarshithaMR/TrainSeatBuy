package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type SeatConfig struct {
	Seats map[string]int `yaml:"seats"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type Config struct {
	SeatConfig   SeatConfig   `yaml:"seats"`
	ServerConfig ServerConfig `yaml:"server"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}
