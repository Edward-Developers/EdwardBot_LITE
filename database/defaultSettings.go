package database

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"github.com/bwmarrin/discordgo"
	"EdwardBot_LITE/structs"
)

func DefaultSettings(g []*discordgo.Guild, session *r.Session) {
	for _, guild := range g {
		r.Table("Settings").Insert(structs.Settings{
			ID: guild.ID,
			PREFIX: ".",
		}).Exec(session)
    }
}