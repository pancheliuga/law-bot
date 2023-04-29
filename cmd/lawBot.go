/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// lawbotCmd represents the lawbot command
var lawbotCmd = &cobra.Command{
	Use:     "lawbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lawbot called")

		lawbot, err := telebot.NewBot(telebot.Settings{
			URL:    "https://api.telegram.org",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Error: please check TELE_TOKEN environment variable - %s", err)
			return
		}

		lawbot.Handle(telebot.OnText, func(c telebot.Context) error {

			log.Print(c.Message().Payload, c.Text())
			payload := c.Message().Payload

			switch payload {
			case "hello":
				err = c.Send(fmt.Sprintf("Hello I'm law_bot %s", appVersion))

			}

			return err
		})

		lawbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(lawbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lawbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lawbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
