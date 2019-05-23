// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/cobra"
)

var parseMode string

var rootCmd = &cobra.Command{
	Use:   "telegram-bot-cli",
	Short: "Could be used to send messages from cli",
	Long: `Command takes two arguments:
A channel
A message 
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		t := viper.GetString("token")
		if "" == t {
			fmt.Println("Provide TELEGRAM_TOKEN env var. You could get it using Telegram BotFather")
			os.Exit(1)
		}

		bot, err := tgbotapi.NewBotAPI(t)
		if err != nil {
			fmt.Printf("Cannot connect to telegram. Check the token is valid. Err: %s\n", err.Error())
			os.Exit(1)
		}

		c, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Printf("Cannot parse first argument as channel id. Check the channel id is intenger. Err: %s\n", err.Error())
			os.Exit(1)
		}

		msg := tgbotapi.NewMessage(c, args[1])


		if tgbotapi.ModeHTML == parseMode {
			msg.ParseMode = tgbotapi.ModeHTML
		} else if tgbotapi.ModeMarkdown == parseMode {
			msg.ParseMode = tgbotapi.ModeMarkdown
		} else if "" == parseMode {
		} else {
			fmt.Printf("Invalid parse mode provided %s. Must be either HTML or Markdown\n", parseMode)
			os.Exit(1)
		}

		if _, err := bot.Send(msg); err != nil {
			fmt.Printf("Failed to send a mesasge to channel. Err: %s\n", err.Error())
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.SetEnvPrefix("telegram")
	viper.AutomaticEnv()

	rootCmd.Flags().StringVar(&parseMode, "parse-mode","", "Set if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.")
}
