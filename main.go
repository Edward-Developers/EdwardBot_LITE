package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
	"EdwardBot_LITE/config"
	"EdwardBot_LITE/database"
	"EdwardBot_LITE/events"
	"EdwardBot_LITE/webpage"
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

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	webpage.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}