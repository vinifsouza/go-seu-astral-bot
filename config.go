package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configs map[string]string

func LoadConfigs() Configs {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("File .env not found, loading environment variables.")
	}

	config := Configs{
		"DISCORD_BOT_TOKEN": os.Getenv("DISCORD_BOT_TOKEN"),
	}

	return config
}
