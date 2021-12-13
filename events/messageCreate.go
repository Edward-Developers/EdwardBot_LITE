package events

import (
	"github.com/bwmarrin/discordgo"
	"EdwardBot_LITE/commands/economy"
	"EdwardBot_LITE/commands/information"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!ping" {
		information.Ping(s, m)
	}
	if m.Content == "!help" {
		information.Help(s, m)
	}
	if m.Content == "!bank" {
		economy.Bank(s, m)
	}
	if m.Content == "!work" {
		economy.Work(s, m)
	}
}