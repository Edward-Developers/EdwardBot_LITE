package events

import (
	"github.com/bwmarrin/discordgo"
	"EdwardBot_LITE/commands"
	"EdwardBot_LITE/config"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!ping" {
		commands.Ping()
		s.ChannelMessageSend(m.ChannelID, "Pong!" + config.Version)
	}
}