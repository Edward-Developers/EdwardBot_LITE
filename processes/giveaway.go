package processes

import (
	"EdwardBot_LITE/modules"
	"github.com/bwmarrin/discordgo"
	"time"
)

func Giveaway(s *discordgo.Session) {
	var ticker = time.NewTicker(30 * time.Second)
	for range ticker.C {
		modules.GiveawayModule(s)
	}
}