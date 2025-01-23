package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type SectionConfig struct {
	Count int `yaml:"count"`
}

type SeatConfig struct {
	A SectionConfig `yaml:"A"`
	B SectionConfig `yaml:"B"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
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
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %v", err)
	}
	fmt.Printf("File content:\n%s\n", string(fileContent))

	// Now decode the content
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}
