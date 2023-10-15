package sounds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func Play(s *discordgo.Session, buffer [][]byte, loopCount *int, guildID, channelID string) (err error) {

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(250 * time.Millisecond)

	// Start speaking.
	vc.Speaking(true)

	// Send the buffer data.
	for {
		*loopCount++
		for _, buff := range buffer {
			vc.OpusSend <- buff
		}
	}
}
