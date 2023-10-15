package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Configs map[string]string

func LoadConfigs() Configs {
	err := godotenv.Load()

	if err != nil {
		color.Yellow("Error reading .env.", err)
		color.Yellow("Loading environment variables.")
	}

	config := Configs{
		"DISCORD_BOT_TOKEN":        os.Getenv("DISCORD_BOT_TOKEN"),
		"APP_LANGUAGE":             os.Getenv("APP_LANGUAGE"),
		"APP_SOUND_PATH":           os.Getenv("APP_SOUND_PATH"),
		"APP_STOP_IMAGES":          os.Getenv("APP_STOP_IMAGES"),
		"APP_STOP_AUTHOR_NAME":     os.Getenv("APP_STOP_AUTHOR_NAME"),
		"APP_STOP_AUTHOR_ICON":     os.Getenv("APP_STOP_AUTHOR_ICON"),
		"APP_STOP_MSG_TITLE":       os.Getenv("APP_STOP_MSG_TITLE"),
		"APP_STOP_MSG_DESCRIPTION": os.Getenv("APP_STOP_MSG_DESCRIPTION"),
	}

	return config
}
