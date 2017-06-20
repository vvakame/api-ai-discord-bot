package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"cloud.google.com/go/compute/metadata"
	"github.com/bwmarrin/discordgo"
	"github.com/k0kubun/pp"
	"github.com/kamalpy/apiai-go"
)

type Config struct {
	discordBOTToken           string
	apiaiDeveloperAccessToken string
}

var config *Config

func main() {
	pp.ColoringEnabled = false

	dg, err := launchBot()
	if err != nil {
		panic(err)
	}
	defer func() {
		dg.Close()
	}()

	fmt.Println("BOT is running...  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func launchBot() (*discordgo.Session, error) {
	var err error
	config, err = getTokens()
	if err != nil {
		return nil, err
	}

	dg, err := discordgo.New("Bot " + config.discordBOTToken)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return nil, err
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return nil, err
	}

	return dg, nil
}

func getTokens() (*Config, error) {
	dToken := os.Getenv("DISCORD_BOT_TOKEN")
	if dToken == "" {
		token, err := metadata.Get("discord-bot-token")
		if err != nil {
			return nil, err
		}

		dToken = token
	}
	if dToken == "" {
		return nil, errors.New("Discord BOT token is empty")
	}

	aToken := os.Getenv("APIAI_DEVELOPER_ACCESS_TOKEN")
	if aToken == "" {
		token, err := metadata.Get("apiai-developer-access-token")
		if err != nil {
			return nil, err
		}
		aToken = token
	}
	if aToken == "" {
		return nil, errors.New("api.ai Developer AccessToken is empty")
	}

	return &Config{
		discordBOTToken:           dToken,
		apiaiDeveloperAccessToken: aToken,
	}, nil
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "mentions me!")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	{
		found := false
		for _, target := range m.Mentions {
			if target.ID == s.State.User.ID {
				found = true
				break
			}
		}

		if !found {
			return
		}
	}

	var content string
	{
		content = m.Content
		if len(m.Mentions) != 0 {
			for _, user := range m.Mentions {
				content = regexp.MustCompile(fmt.Sprintf("<@%s>", user.ID)).ReplaceAllString(content, "")
			}
		}
		content = strings.TrimSpace(content)
	}

	fmt.Println("query", content)

	ai := &apiaigo.APIAI{
		AuthToken: config.apiaiDeveloperAccessToken,
		Language:  "ja",
		SessionID: fmt.Sprintf("c%v", m.ChannelID),
		Version:   "20170611",
	}
	queryResp, err := ai.SendText(content)
	if err != nil {
		fmt.Printf("ai.SendText: %v", err.Error())
		fmt.Errorf("ai.SendText: %v", err.Error())
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, queryResp.Result.Fulfillment.Messages[0].Speech)
	if err != nil {
		fmt.Printf("s.ChannelMessageSend: %v", err.Error())
		fmt.Errorf("s.ChannelMessageSend: %v", err.Error())
		return
	}

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Printf("s.State.Channel: %v", err.Error())
		fmt.Errorf("s.State.Channel: %v", err.Error())
		return
	}

	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Printf("s.State.Guild: %v", err.Error())
		fmt.Errorf("s.State.Guild: %v", err.Error())
		return
	}

	pp.Println(guild, channel, m.Message, queryResp)
}
