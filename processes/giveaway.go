package processes

import (
	"fmt"
	"time"
	"EdwardBot_LITE/events"
)

func Giveaway() {
	var guilds = events.Guilds
	var tick = time.Tick(60000 * time.Millisecond)
	for range tick {
		for _, guild := range guilds {
			fmt.Println(guild.ID)
		}
	}
}