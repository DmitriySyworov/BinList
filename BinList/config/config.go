package config

import (
	"os"

	"github.com/fatih/color"
)

type Config struct {
	MasterKey string
	AccessKey string
}

func NewConfig() *Config {
	keyMaster := os.Getenv("Master")
	if keyMaster == "" {
		panic(color.RedString("Переменная окружения MasterKey  не задана!"))
	}
	keyAccess := os.Getenv("Access")
	if keyAccess == "" {
		panic(color.RedString("Переменная окружения AccessKey не задана!"))
	}
	return &Config{
		MasterKey: keyMaster,
		AccessKey: keyAccess,
	}
}
