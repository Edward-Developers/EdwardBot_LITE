package economy

import (
	"EdwardBot_LITE/database"
	"EdwardBot_LITE/structs"
	"fmt"
	"math/rand"
	"strconv"
	"github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func Work(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {//co sie tu ... ale działa :P
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
	economies, _ = r.Table("Economies").Get(m.Author.ID).Run(database.Session)
	economies.One(&row)
	data, _ := row.(map[string]interface{})
	var coin = data["COIN"]
	var bank = data["BANK"]
	var coinString = fmt.Sprint(coin)
	var bankString = fmt.Sprint(bank)
	var min = 10
    var max = 70
	coinInt, _ := strconv.ParseInt(coinString, 0, 64)
	bankInt, _ := strconv.ParseInt(bankString, 0, 64)
	var randomInt = rand.Intn(max - min) + min
	var coinUpdate = coinInt + int64(randomInt)
	r.Table("Economies").Get(m.Author.ID).Update(structs.EconomyUser{
		ID: m.Author.ID,
		COIN: coinUpdate,
		BANK: bankInt,
	}).Exec(database.Session)
	economies, _ = r.Table("Economies").Get(m.Author.ID).Run(database.Session)
	economies.One(&row)
	dataAfterUpdate, _ := row.(map[string]interface{})
	var coinAfterUpdate = dataAfterUpdate["COIN"]
	var coinStringAfterUpdate = fmt.Sprint(coinAfterUpdate)
	s.ChannelMessageSend(m.ChannelID, "Work " + coinStringAfterUpdate)
	defer economies.Close()
}