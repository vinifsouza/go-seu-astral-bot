package commands

import (
	"go-seu-astral-bot/sounds"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func Start(s *discordgo.Session, m *discordgo.MessageCreate, buffer [][]byte, loopCount *int) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		color.Red("Error s.State.Channel: %v", err)
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		color.Red("Error s.State.Guild: %v", err)
		return
	}

	// Look for the message sender in that guild's current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			err = sounds.Play(s, buffer, loopCount, g.ID, vs.ChannelID)
			if err != nil {
				color.Red("Error playing sound: %v", err)
			}

			return
		}
	}
}
