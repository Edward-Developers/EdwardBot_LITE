package information

import (
	"EdwardBot_LITE/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong! "+config.Version)
	if err != nil {
		return
	}
	fmt.Println(g)
}
