package modules

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/snowflake/v2"
	"golang.org/x/net/context"
	"os"
	"strconv"
)

var Lavalink disgolink.Client

func LavalinkModule(s *discordgo.State) disgolink.Client {
	parseInt, _ := strconv.ParseInt(s.User.ID, 10, 64)
	var userID = snowflake.ID(parseInt)

	lavalinkClient := disgolink.New(userID,
		disgolink.WithListenerFunc(func(p disgolink.Player, event lavalink.Event) {
			fmt.Printf("EVENT: %T\n", event)
		}),
	)

	_, err := lavalinkClient.AddNode(context.TODO(), disgolink.NodeConfig{
		Name:     "main",
		Address:  "138.2.132.153:7771",
		Password: os.Getenv("LAVALINK_PASSWORD"),
		Secure:   false,
	})
	if err != nil {
		return nil
	}
	return lavalinkClient
}