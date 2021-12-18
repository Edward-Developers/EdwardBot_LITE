package events

import (
	"github.com/bwmarrin/discordgo"
	"EdwardBot_LITE/database"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"EdwardBot_LITE/commands/economy"
	"EdwardBot_LITE/commands/information"
	"EdwardBot_LITE/commands/settings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	settingsP, _ := r.Table("Settings").Get(m.GuildID).Run(database.Session)
	var row interface{}
	err := settingsP.One(&row)
	if err == r.ErrEmptyResult {
		return
	}
	data, _ := row.(map[string]interface{})
	var prefix = data["PREFIX"].(string)
	if m.Content ==  prefix + "ping" {
		information.Ping(s, m, Guilds)
	}
	if m.Content == prefix + "help" {
		information.Help(s, m, Guilds)
	}
	if m.Content == prefix + "bank" {
		economy.Bank(s, m, Guilds)
	}
	if m.Content == prefix + "work" {
		economy.Work(s, m, Guilds)
	}
	if m.Content == prefix + "prefix" {
		settings.Prefix(s, m, Guilds)
	}
	defer settingsP.Close()
}