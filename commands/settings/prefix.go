package settings

import (
	"github.com/bwmarrin/discordgo"
)

func Prefix(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {
	s.ChannelMessageSend(m.ChannelID, "Prefix!")
}