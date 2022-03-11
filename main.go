package main

import (
	"fmt"
	"os"

	teams "github.com/atc0005/go-teams-notify/v2"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	WebHookUrl        string `short:"w" long:"webhook-url" description:"Microsoft Teams webhook URL" env:"WEBHOOK_URL" required:"true"`
	CardTitle         string `short:"t" long:"title" description:"Title of the card" required:"true"`
	CardDescription   string `short:"d" long:"description" description:"The description of the card (supports HTML)" required:"true"`
	Status            string `short:"s" long:"status" description:"Status (i.e. success, canceled, failed)" env:"CI_JOB_STATUS" required:"true"`
	ActionUrl         string `short:"a" long:"action-url" description:"The action URL for the button on the card" env:"CI_JOB_URL" required:"false"`
	ActionButtonTitle string `short:"b" long:"action-button-title" description:"The title of the action button" default:"View in GitLab" required:"false"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)
	}

	mstClient := teams.NewClient()

	msgCard := teams.NewMessageCard()
	msgCard.Title = opts.CardTitle
	msgCard.Text = opts.CardDescription

	switch opts.Status {
	case "success":
		msgCard.ThemeColor = "#007500"
	case "canceled":
		msgCard.ThemeColor = "#808080"
	case "failed":
		msgCard.ThemeColor = "#FF0000"
	default:
		die(fmt.Sprintf("Unsupported status '%s'", opts.Status))
	}

	if opts.ActionUrl != "" {
		pa, err := teams.NewMessageCardPotentialAction(teams.PotentialActionOpenURIType, opts.ActionButtonTitle)

		if err != nil {
			die("Unable to create new action: " + err.Error())
		}

		pa.MessageCardPotentialActionOpenURI.Targets = []teams.MessageCardPotentialActionOpenURITarget{{OS: "default", URI: opts.ActionUrl}}

		if err := msgCard.AddPotentialAction(pa); err != nil {
			die("Unable to add action to message card: " + err.Error())
		}
	}

	err = mstClient.Send(opts.WebHookUrl, msgCard)

	if err != nil {
		die(err.Error())
	}
}

func die(err string) {
	fmt.Printf("error: %s", err)
	os.Exit(1)
}
