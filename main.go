package main

import (
	"EdwardBot_LITE/config"
	"EdwardBot_LITE/database"
	"EdwardBot_LITE/events"
	"EdwardBot_LITE/modules"
	"EdwardBot_LITE/processes"
	"EdwardBot_LITE/webpage"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/snowflake/v2"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var Connect = database.Connect()
	database.Session = Connect
	database.CreateTables(Connect)

	dg, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(events.GuildCreate)
	dg.AddHandler(events.GuildDelete)
	dg.AddHandler(events.MessageCreate)
	dg.AddHandler(events.MessageDelete)
	dg.AddHandler(events.Ready)

	dg.AddHandler(func(s *discordgo.Session, v *discordgo.VoiceServerUpdate) {
		guildID, _ := snowflake.Parse(v.GuildID)
		modules.Lavalink.OnVoiceServerUpdate(context.TODO(), guildID, v.Token, v.Endpoint)
	})

	dg.AddHandler(func(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
		if v.UserID != s.State.User.ID {
			return
		}
		guildID, _ := snowflake.Parse(v.GuildID)
		var channelID *snowflake.ID
		if v.ChannelID != "" {
			parsed, _ := snowflake.Parse(v.ChannelID)
			channelID = &parsed
		}
		modules.Lavalink.OnVoiceStateUpdate(context.TODO(), guildID, channelID, v.SessionID)
	})

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	go processes.Giveaway(dg)

	var lava = modules.LavalinkModule(dg.State)
	modules.Lavalink = lava
	webpage.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	err = dg.Close()
	if err != nil {
		return
	}
}