package events

import (
	"github.com/bwmarrin/discordgo"
	"EdwardBot_LITE/commands"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!ping" {
		commands.Ping(s, m)
	}
	if m.Content == "!help" {
		commands.Help(s, m)
	}
}