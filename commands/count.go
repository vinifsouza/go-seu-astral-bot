package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Count(dcSession *discordgo.Session, dcMessage *discordgo.MessageCreate, loopCount *int) {
	loopCountStr := strconv.Itoa(*loopCount)
	if *loopCount == 1 {
		dcSession.ChannelMessageSend(dcMessage.ChannelID, "A música já foi tocada "+loopCountStr+" vez")
	} else {
		dcSession.ChannelMessageSend(dcMessage.ChannelID, "A música já foi tocada "+loopCountStr+" vezes")
	}
}
