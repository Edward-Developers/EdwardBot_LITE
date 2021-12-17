package economy

import (
	"github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"EdwardBot_LITE/database"
	"EdwardBot_LITE/structs"
)

func Bank(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {
	var guildID = m.GuildID
	r.Table("Economies").Insert(structs.EconomyUser{
		ID: m.Author.ID,
		COIN: 0,
		BANK: 0,
	}).Exec(database.Session)
	s.ChannelMessageSend(m.ChannelID, "Bank "+ guildID)
}