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

		if args[1] == "help" {
			embed := discordgo.MessageEmbed{
				Title: "`!go tip` for golang tips\n`!go joke` for jokes\n`!go fact` for facts\n`!go define your_word` for word definitions\n",
			}
			s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		}
		if args[1] == "tip" {
			//author := discordgo.MessageEmbedAuthor{
			//	Name: "Mahady Hasan Pial",
			//	URL:  "https://github.com/techno-stupid",
			//}
			embed := discordgo.MessageEmbed{
				Title: helpers.RandomTip(),
				//Author: &author,
			}
			s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		}
		if args[1] == "joke" {
			joke, jokeError := helpers.GetJoke()
			if jokeError != nil {
				joke = "No joke found"
			}
			s.ChannelMessageSend(m.ChannelID, joke)
		}
		if args[1] == "fact" {
			fact, factError := helpers.GetFact()
			if factError != nil {
				fact = "No fact found"
			}
			s.ChannelMessageSend(m.ChannelID, fact)
		}
		if args[1] == "define" {
			definition, definitionError := helpers.GetDefinition(args[2])
			fmt.Println(definition)
			if definitionError == nil {
				s.ChannelMessageSend(m.ChannelID, definition)
			}
		}
	})

	ses.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
		fmt.Println(r.Emoji.Name)
		if r.Emoji.Name == "🤘" {
			s.GuildMemberRoleAdd(r.GuildID, r.UserID, "1132173518369468436")
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("%v has been added to %v", r.UserID, r.Emoji.Name))
		}
	})

	ses.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
		fmt.Println(r.Emoji.Name)
		if r.Emoji.Name == "🤘" {
			s.GuildMemberRoleRemove(r.GuildID, r.UserID, "1132173518369468436")
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("%v has been removed to %v", r.UserID, r.Emoji.Name))
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
