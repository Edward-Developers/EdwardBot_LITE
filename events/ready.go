package events

import (
	"EdwardBot_LITE/database"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var Guilds []*discordgo.Guild

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	Guilds = r.Guilds
	fmt.Println("Bot is now running!")
	database.DefaultSettings(Guilds, database.Session)
}