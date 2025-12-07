package config

import "os"

type Config struct {
	MasterKey string
	AccessKey string
}

func NewConfig() *Config {
	keyMaster := os.Getenv("MasterKey")
	if keyMaster == "" {
		panic("Переменная окружения MasterKey  не задана!")
	}
	keyAccess := os.Getenv("AccessKey")
	if keyAccess == "" {
		panic("Переменная окружения AccessKey не задана!")
	}
	return &Config{
		MasterKey: keyMaster,
		AccessKey: keyAccess,
	}
}
