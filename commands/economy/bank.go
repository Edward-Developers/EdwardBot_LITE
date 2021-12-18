package economy

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"EdwardBot_LITE/database"
	"EdwardBot_LITE/structs"
)

func Bank(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {
	economies, _ := r.Table("Economies").Get(m.Author.ID).Run(database.Session)
	var row interface{}
	err := economies.One(&row)
	if err == r.ErrEmptyResult {
		r.Table("Economies").Insert(structs.EconomyUser{
			ID: m.Author.ID,
			COIN: 0,
			BANK: 0,
		}).Exec(database.Session)
	}
	data, _ := row.(map[string]interface{})
	var coin = data["COIN"]
	var bank = data["BANK"]
	s.ChannelMessageSend(m.ChannelID, "Bank\nCoin: " + fmt.Sprint(coin) + "\nBank: " + fmt.Sprint(bank))
	defer economies.Close()
}