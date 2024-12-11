package slack

import (
	"github.com/slack-go/slack"
)

type Client struct {
	api *slack.Client
}

func NewClient(botToken string) *Client {
	return &Client{
		api: slack.New(botToken),
	}
}
