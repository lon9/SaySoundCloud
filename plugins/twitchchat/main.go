package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/gempir/go-twitch-irc"
)

type body struct {
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

func main() {
	var (
		nickname    string
		channel     string
		twitchToken string
		endpoint    string
		accessToken string
	)
	flag.StringVar(&nickname, "n", "SaySoundBot", "Nickname")
	flag.StringVar(&channel, "c", "", "Channel")
	flag.StringVar(&twitchToken, "t", "", "OAuth token")
	flag.StringVar(&endpoint, "u", "http://localhost:3001/v1/apps/18/cmd", "server endpoint")
	flag.StringVar(&accessToken, "a", "", "SaySoundCloud's access token")
	flag.Parse()

	if channel == "" || twitchToken == "" || accessToken == "" {
		panic(errors.New("channel, twitch token and access token options must be set"))
	}
	client := twitch.NewClient(nickname, twitchToken)
	client.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		cmd := strings.TrimPrefix(msg.Message, "!")
		if len(cmd) == 0 {
			return
		}
		b := &body{
			Name:        cmd,
			AccessToken: accessToken,
		}
		bt, err := json.Marshal(b)
		if err != nil {
			panic(err)
		}
		resp, err := http.Post(endpoint, "application/json", bytes.NewReader(bt))
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
	})
	client.Join(channel)
	if err := client.Connect(); err != nil {
		panic(err)
	}
}
