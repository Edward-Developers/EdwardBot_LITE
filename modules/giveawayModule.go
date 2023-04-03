package modules

import (
	"EdwardBot_LITE/events"
	"fmt"
)

func GiveawayModule() {
	var guilds = events.Guilds
	for _, guild := range guilds {
		fmt.Println(guild.ID)
	}
}
