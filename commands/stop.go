package commands

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Stop(dcSession *discordgo.Session, dcMessage *discordgo.MessageCreate, configs map[string]string) {
	rand.Seed(time.Now().UnixNano())

	images := strings.Split(configs["APP_STOP_IMAGES"], ",")
	randomIndex := rand.Intn(len(images))

	image := discordgo.MessageEmbedImage{
		URL: images[randomIndex],
	}

	author := discordgo.MessageEmbedAuthor{
		Name:    configs["APP_STOP_AUTHOR_NAME"],
		IconURL: configs["APP_STOP_AUTHOR_ICON"],
	}

	dcSession.ChannelMessageSendEmbed(dcMessage.ChannelID, &discordgo.MessageEmbed{
		Author:      &author,
		Image:       &image,
		Title:       configs["APP_STOP_MSG_TITLE"],
		Description: configs["APP_STOP_MSG_DESCRIPTION"],
	})
}
