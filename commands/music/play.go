package music

import (
	"github.com/bwmarrin/discordgo"
)

func Play(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Play")
	if err != nil {
		return
	}
}
