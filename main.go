package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"go-discord-bot/helpers"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const command = "!go"

func main() {
	ses, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	ses.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Split(m.Content, " ")

		if args[0] != command {
			return
		}

		if args[1] == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world!")
		}

		if args[1] == "tip" {
			s.ChannelMessageSend(m.ChannelID, helpers.RandomTip())
		}
	})

	ses.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = ses.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer ses.Close()

	fmt.Println("The bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
}
