package events

import (
	"EdwardBot_LITE/database"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Bot is now running!")
	database.DefaultSettings(r.Guilds, database.Session)
}