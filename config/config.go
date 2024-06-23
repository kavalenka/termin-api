package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type DBConfig struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
}

func NewDB() *DBConfig {
	configFile, err := os.ReadFile("./config/database.yml")

	if err != nil {
		log.Fatalf("Failed to read DB config file: %s", err)
	}

	var c DBConfig

	err = yaml.Unmarshal(configFile, &c)

	if err != nil {
		log.Fatalf("Failed to decode DB config file: %s", err)
	}

	return &c
}
