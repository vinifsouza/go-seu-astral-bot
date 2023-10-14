package commands

import (
	"fmt"
	"go-seu-astral-bot/sounds"

	"github.com/bwmarrin/discordgo"
)

func Start(s *discordgo.Session, m *discordgo.MessageCreate, buffer [][]byte) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Error s.State.Channel: ", err)
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		fmt.Println("Error s.State.Guild: ", err)
		return
	}

	// Look for the message sender in that guild's current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			err = sounds.Play(s, buffer, g.ID, vs.ChannelID)
			if err != nil {
				fmt.Println("Error playing sound: ", err)
			}

			return
		}
	}
}
