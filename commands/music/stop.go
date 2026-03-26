package music

import (
	"EdwardBot_LITE/modules"
	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/snowflake/v2"
	"golang.org/x/net/context"
)

func Stop(s *discordgo.Session, m *discordgo.MessageCreate) {
	guildIDNode, _ := snowflake.Parse(m.GuildID)
	player := modules.Lavalink.Player(guildIDNode)

	err := player.Destroy(context.TODO())
	if err != nil {
		return
	}

	err = s.ChannelVoiceJoinManual(m.GuildID, "", false, false)
	if err != nil {
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, "Stopped playing and left the channel.")
	if err != nil {
		return
	}
}