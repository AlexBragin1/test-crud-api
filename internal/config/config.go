package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port string `yaml:"port" env-default:"8080"`
	DB   `yaml:"db"`
}

type DB struct {
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
	Password string `yaml:"password"`
}

func LoadDB() *Config {
	os.Setenv("POSTGRES_STORAGE", "/home/kroot/My_program/test-crud-api/config/config.yaml")
	configPath := os.Getenv("POSTGRES_STORAGE")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
