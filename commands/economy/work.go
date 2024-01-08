package economy

import (
	"EdwardBot_LITE/database"
	"EdwardBot_LITE/structs"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"math/rand"
	"strconv"
)

func Work(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {
	economies, _ := r.Table("Economies").Get(m.Author.ID).Run(database.Session)
	var row interface{}
	err := economies.One(&row)
	if errors.Is(err, r.ErrEmptyResult) {
		err := r.Table("Economies").Insert(structs.EconomyUser{
			ID:   m.Author.ID,
			COIN: 0,
			BANK: 0,
		}).Exec(database.Session)
		if err != nil {
			return
		}
	}
	economies, _ = r.Table("Economies").Get(m.Author.ID).Run(database.Session)
	err = economies.One(&row)
	if err != nil {
		return
	}
	data, _ := row.(map[string]interface{})
	var coin = data["COIN"]
	var bank = data["BANK"]
	var coinString = fmt.Sprint(coin)
	var bankString = fmt.Sprint(bank)
	var min = 10
	var max = 70
	coinInt, _ := strconv.ParseInt(coinString, 0, 64)
	bankInt, _ := strconv.ParseInt(bankString, 0, 64)
	var randomInt = rand.Intn(max-min) + min
	var coinUpdate = coinInt + int64(randomInt)
	err = r.Table("Economies").Get(m.Author.ID).Update(structs.EconomyUserUpdate{
		COIN: coinUpdate,
		BANK: bankInt,
	}).Exec(database.Session)
	if err != nil {
		return
	}
	economies, _ = r.Table("Economies").Get(m.Author.ID).Run(database.Session)
	err = economies.One(&row)
	if err != nil {
		return
	}
	dataAfterUpdate, _ := row.(map[string]interface{})
	var coinAfterUpdate = dataAfterUpdate["COIN"]
	var coinStringAfterUpdate = fmt.Sprint(coinAfterUpdate)
	_, err = s.ChannelMessageSend(m.ChannelID, "Work "+coinStringAfterUpdate)
	if err != nil {
		return
	}
	defer func(economies *r.Cursor) {
		err := economies.Close()
		if err != nil {
			return
		}
	}(economies)
}
