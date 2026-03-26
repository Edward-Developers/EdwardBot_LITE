package music

import (
	"EdwardBot_LITE/modules"
	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/snowflake/v2"
	"golang.org/x/net/context"
)

func Play(s *discordgo.Session, m *discordgo.MessageCreate) {
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		return
	}

	var voiceChannelID string
	for _, vs := range guild.VoiceStates {
		if vs.UserID == m.Author.ID {
			voiceChannelID = vs.ChannelID
			break
		}
	}

	if voiceChannelID == "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Join a voice channel first!")
		return
	}

	err = s.ChannelVoiceJoinManual(m.GuildID, voiceChannelID, false, true)
	if err != nil {
		return
	}

	guildIDNode, _ := snowflake.Parse(m.GuildID)
	player := modules.Lavalink.Player(guildIDNode)

	query := "ytsearch:Rick Astley - Never Gonna Give You Up"

	modules.Lavalink.BestNode().LoadTracksHandler(context.TODO(), query, disgolink.NewResultHandler(
		func(track lavalink.Track) {
			err := player.Update(context.TODO(), lavalink.WithTrack(track))
			if err == nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Playing: "+track.Info.Title)
			}
		},
		func(playlist lavalink.Playlist) {},
		func(tracks []lavalink.Track) {
			err := player.Update(context.TODO(), lavalink.WithTrack(tracks[0]))
			if err == nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Playing: "+tracks[0].Info.Title)
			}
		},
		func() {},
		func(err error) {},
	))
}