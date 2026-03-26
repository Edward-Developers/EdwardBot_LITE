package information

import (
	"EdwardBot_LITE/config"
	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong! "+config.Version)
	if err != nil {
		return
	}
}