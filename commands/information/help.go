package information

import (
	"github.com/bwmarrin/discordgo"
)
func Help(s *discordgo.Session, m *discordgo.MessageCreate, g []*discordgo.Guild) {
	var fields []*discordgo.MessageEmbedField
	var field1 = &discordgo.MessageEmbedField{
		Name:   "Test",
		Value:  "testowanie",
		Inline: true,
	}
	var field2 = &discordgo.MessageEmbedField{
		Name:   "Test",
		Value:  "testowanie",
		Inline: true,
	}
	fields = append(fields, field1)
	fields = append(fields, field2)
	var embed = &discordgo.MessageEmbed{
		Type: "rich",
		Title: "Test",
		Description: "Test",
		Fields: fields,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}