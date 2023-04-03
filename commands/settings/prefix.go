package settings

import (
	"github.com/bwmarrin/discordgo"
)

func Prefix(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Prefix!")
	if err != nil {
		return
	}
}
