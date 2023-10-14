package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs map[string]string

func LoadConfigs() Configs {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading envs: ", err)
	}

	config := Configs{
		"DISCORD_BOT_TOKEN": os.Getenv("DISCORD_BOT_TOKEN"),
	}

	return config
}
