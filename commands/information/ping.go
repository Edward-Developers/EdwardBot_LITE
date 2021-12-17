package information

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"EdwardBot_LITE/config"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {
	s.ChannelMessageSend(m.ChannelID, "Pong! " + config.Version)
	fmt.Println(g)
}