package economy

import (
	"github.com/bwmarrin/discordgo"
)

func Shop(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Shop ")
	if err != nil {
		return
	}
}
