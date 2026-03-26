package economy

import (
	"EdwardBot_LITE/database"
	"EdwardBot_LITE/structs"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"math/rand"
)

func Work(s *discordgo.Session, m *discordgo.MessageCreate) {
	economies, err := r.Table("Economies").Get(m.Author.ID).Run(database.Session)
	if err != nil {
		return
	}
	defer economies.Close()

	var user structs.EconomyUser
	err = economies.One(&user)

	if errors.Is(err, r.ErrEmptyResult) {
		user = structs.EconomyUser{
			ID:   m.Author.ID,
			COIN: 0,
			BANK: 0,
		}
		err := r.Table("Economies").Insert(user).Exec(database.Session)
		if err != nil {
			return
		}
	}

	var min = 10
	var max = 70
	var randomInt = rand.Intn(max-min) + min

	user.COIN += int64(randomInt)

	err = r.Table("Economies").Get(m.Author.ID).Update(structs.EconomyUserUpdate{
		COIN: user.COIN,
		BANK: user.BANK,
	}).Exec(database.Session)
	if err != nil {
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, "Work "+fmt.Sprint(user.COIN))
	if err != nil {
		return
	}
}