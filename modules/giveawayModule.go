package modules

import (
	"fmt"
	"EdwardBot_LITE/events"
)

func GiveawayModule() {
	var guilds = events.Guilds
	for _, guild := range guilds {
		fmt.Println(guild.ID)
	}
}