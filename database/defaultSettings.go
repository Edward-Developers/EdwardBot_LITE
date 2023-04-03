package database

import (
	"EdwardBot_LITE/structs"
	"github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func DefaultSettings(g []*discordgo.Guild, session *r.Session) {
	for _, guild := range g {
		err := r.Table("Settings").Insert(structs.Settings{
			ID:     guild.ID,
			PREFIX: ".",
		}).Exec(session)
		if err != nil {
			return
		}
	}
}
