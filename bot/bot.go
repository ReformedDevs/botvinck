package bot

import (
	"github.com/nlopes/slack"
)

// Bot interacts with the Slack API, responding to messages as required.
type Bot struct {
	client *slack.Client
}

// New creates and initializes a new bot instance.
func New(slackToken string) *Bot {
	c := slack.New(slackToken)
	return &Bot{
		client: c,
	}
}

// Close shuts down the bot.
func (b *Bot) Close() {
	//...
}
