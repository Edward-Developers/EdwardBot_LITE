package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/disgolink"
	"github.com/disgoorg/snowflake/v2"
	"golang.org/x/net/context"
	"os"
	_ "os"
	"strconv"
)

func LavalinkModule(s *discordgo.State) {
	parseInt, _ := strconv.ParseInt(s.User.ID, 10, 64)
	var userID = snowflake.ID(parseInt)
	lavalinkClient := disgolink.New(userID)
	_, err := lavalinkClient.AddNode(context.TODO(), disgolink.NodeConfig{
		Name:      "main",
		Address:   "130.61.95.106:8080",
		Password:  os.Getenv("LAVALINK_PASSWORD"),
		Secure:    false,
		SessionID: "",
	})
	if err != nil {
		return
	}
}
