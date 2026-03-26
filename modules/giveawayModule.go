package modules

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func GiveawayModule(s *discordgo.Session) {
	var guilds = s.State.Guilds
	for _, guild := range guilds {
		fmt.Println(guild.ID)
	}
}