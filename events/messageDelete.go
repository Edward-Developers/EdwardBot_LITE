package events

import (
	"github.com/bwmarrin/discordgo"
)

var DeleteMessages []interface{}

func MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	
}