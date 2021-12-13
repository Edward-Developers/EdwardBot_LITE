package economy

import (
	"github.com/bwmarrin/discordgo"
)

func Work(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Work")
}