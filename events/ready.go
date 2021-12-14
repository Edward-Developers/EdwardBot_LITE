package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	for _, guild := range r.Guilds {
		fmt.Println(guild.ID)
    }
	fmt.Println("Bot is now running!")
}