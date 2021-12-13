package economy

import (
	"github.com/bwmarrin/discordgo"
)

func Bank(s *discordgo.Session, m *discordgo.MessageCreate) {
	var guildID = m.GuildID
	s.ChannelMessageSend(m.ChannelID, "Bank "+ guildID)
}