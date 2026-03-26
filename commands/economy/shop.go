package economy

import (
	"github.com/bwmarrin/discordgo"
)

func Shop(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Shop ")
	if err != nil {
		return
	}
}