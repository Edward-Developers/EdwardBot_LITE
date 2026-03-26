package economy

import (
	"EdwardBot_LITE/database"
	"EdwardBot_LITE/structs"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func Bank(s *discordgo.Session, m *discordgo.MessageCreate) {
	economies, err := r.Table("Economies").Get(m.Author.ID).Run(database.Session)
	if err != nil {
		return
	}
	defer economies.Close()

	var row interface{}
	err = economies.One(&row)

	var coin interface{} = 0
	var bank interface{} = 0

	if errors.Is(err, r.ErrEmptyResult) {
		err := r.Table("Economies").Insert(structs.EconomyUser{
			ID:   m.Author.ID,
			COIN: 0,
			BANK: 0,
		}).Exec(database.Session)
		if err != nil {
			return
		}
	} else if err == nil {
		data, _ := row.(map[string]interface{})
		coin = data["COIN"]
		bank = data["BANK"]
	}

	_, err = s.ChannelMessageSend(m.ChannelID, "Bank\nCoin: "+fmt.Sprint(coin)+"\nBank: "+fmt.Sprint(bank))
	if err != nil {
		return
	}
}