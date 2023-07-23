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

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	processes.Giveaway()
	modules.LavalinkModule(dg.State)
	webpage.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	err = dg.Close()
	if err != nil {
		return
	}
}
