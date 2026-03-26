package settings

import (
	"github.com/bwmarrin/discordgo"
)

func Prefix(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Prefix!")
	if err != nil {
		return
	}
}