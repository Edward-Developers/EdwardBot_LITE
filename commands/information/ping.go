package information

import (
	"github.com/bwmarrin/discordgo"
	"EdwardBot_LITE/config"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong! " + config.Version)
}