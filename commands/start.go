package commands

import (
	"go-seu-astral-bot/sounds"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func Start(dcSession *discordgo.Session, dcMessage *discordgo.MessageCreate, buffer [][]byte, loopCount *int) {
	channel, err := dcSession.State.Channel(dcMessage.ChannelID)
	if err != nil {
		color.Red("Error s.State.Channel: %v", err)
		return
	}

	// Find the guild for that channel.
	guild, err := dcSession.State.Guild(channel.GuildID)
	if err != nil {
		color.Red("Error s.State.Guild: %v", err)
		return
	}

	// Look for the message sender in that guild's current voice states.
	for _, voiceState := range guild.VoiceStates {
		if voiceState.UserID == dcMessage.Author.ID {
			err = sounds.Play(dcSession, buffer, loopCount, guild.ID, voiceState.ChannelID)
			if err != nil {
				color.Red("Error playing sound: %v", err)
			}

			return
		}
	}
}
