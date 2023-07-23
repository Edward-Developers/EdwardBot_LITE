package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/disgolink"
	"github.com/disgoorg/snowflake/v2"
	"golang.org/x/net/context"
	"strconv"
)

func LavalinkModule(s *discordgo.State) {
	parseInt, _ := strconv.ParseInt(s.User.ID, 10, 64)
	var userID = snowflake.ID(parseInt)
	lavalinkClient := disgolink.New(userID)
	_, _ = lavalinkClient.AddNode(context.TODO(), disgolink.NodeConfig{
		Name:      "",
		Address:   "",
		Password:  "",
		Secure:    false,
		SessionID: "",
	})
}
