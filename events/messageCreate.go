package events

import (
	"EdwardBot_LITE/commands/economy"
	"EdwardBot_LITE/commands/information"
	"EdwardBot_LITE/commands/music"
	"EdwardBot_LITE/commands/settings"
	"EdwardBot_LITE/database"
	"errors"
	"github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	settingsP, err := r.Table("Settings").Get(m.GuildID).Run(database.Session)
	if err != nil {
		return
	}
	defer settingsP.Close()

	var row interface{}
	err = settingsP.One(&row)
	if errors.Is(err, r.ErrEmptyResult) {
		return
	}

	data, _ := row.(map[string]interface{})
	var prefix = data["PREFIX"].(string)

	if m.Content == prefix+"ping" {
		information.Ping(s, m)
	}
	if m.Content == prefix+"help" {
		information.Help(s, m)
	}
	if m.Content == prefix+"bank" {
		economy.Bank(s, m)
	}
	if m.Content == prefix+"shop" {
		economy.Shop(s, m)
	}
	if m.Content == prefix+"work" {
		economy.Work(s, m)
	}
	if m.Content == prefix+"prefix" {
		settings.Prefix(s, m)
	}
	if m.Content == prefix+"play" {
		music.Play(s, m)
	}
	if m.Content == prefix+"stop" {
		music.Stop(s, m)
	}
}