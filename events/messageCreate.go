package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"EdwardBot_LITE/database"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"EdwardBot_LITE/commands/economy"
	"EdwardBot_LITE/commands/information"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	settings, _ := r.Table("Settings").Get(m.GuildID).Run(database.Session)
	fmt.Println(settings) // jak to odczytać?
	if m.Content == "!ping" {
		information.Ping(s, m, Guilds)
	}
	if m.Content == "!help" {
		information.Help(s, m, Guilds)
	}
	if m.Content == "!bank" {
		economy.Bank(s, m, Guilds)
	}
	if m.Content == "!work" {
		economy.Work(s, m, Guilds)
	}
}