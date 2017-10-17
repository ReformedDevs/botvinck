package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ReformedDevs/botvinck/bot"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "botvinck"
	app.Usage = "Slack bot with a distinctly reformed personality"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "slack-token",
			EnvVar: "SLACK_TOKEN",
			Usage:  "API token for Slack",
		},
	}
	app.Action = func(c *cli.Context) error {

		// Create the bot and run it
		b := bot.New(c.String("slack-token"))
		defer b.Close()

		// Wait for SIGINT or SIGTERM
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: %s\n", err.Error())
	}
}
