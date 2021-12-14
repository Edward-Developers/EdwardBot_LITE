package economy

import (
	"github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"EdwardBot_LITE/database"
)

func Bank(s *discordgo.Session, m *discordgo.MessageCreate) {
	var guildID = m.GuildID
	r.Table("Economies").Insert(map[string]string{
		"id": m.Author.ID,
		"coin": "0",
		"bank": "0",
	}).Exec(database.Session)
	s.ChannelMessageSend(m.ChannelID, "Bank "+ guildID)
}