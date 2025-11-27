package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	keyEnv := os.Getenv("KEY")
	if keyEnv == "" {
		panic("Переменная окружения KEY  не задана!")
	}
	return &Config{
		Key: keyEnv,
	}
}
