package events

import (
	"EdwardBot_LITE/commands/economy"
	"EdwardBot_LITE/commands/information"
	"EdwardBot_LITE/commands/music"
	"EdwardBot_LITE/commands/settings"
	"EdwardBot_LITE/database"
	"github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
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
	if m.Content == prefix+"ping" {
		information.Ping(s, m, Guilds)
	}
	if m.Content == prefix+"help" {
		information.Help(s, m, Guilds)
	}
	if m.Content == prefix+"bank" {
		economy.Bank(s, m, Guilds)
	}
	if m.Content == prefix+"shop" {
		economy.Shop(s, m, Guilds)
	}
	if m.Content == prefix+"work" {
		economy.Work(s, m, Guilds)
	}
	if m.Content == prefix+"prefix" {
		settings.Prefix(s, m, Guilds)
	}
	if m.Content == prefix+"music" {
		music.Play(s, m, Guilds)
	}
	defer func(settingsP *r.Cursor) {
		err := settingsP.Close()
		if err != nil {
			return
		}
	}(settingsP)
}
