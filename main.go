package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	TOKEN_BOT := os.Getenv("TOKEN_LOUROJOSE")
	session, err := discordgo.New("Bot " + TOKEN_BOT)
	if err != nil {
		log.Fatal(err)
	}
	session.AddHandler(func(session *discordgo.Session, msg *discordgo.MessageCreate) {
		/* 		if msg.Author.ID = session.State.User.ID {
			return
		} */
		if msg.Content == "buzz" {
			session.ChannelMessageSend(msg.ChannelID, "fizz")
		}
	})
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	fmt.Println("The Bot Is Online...")
	signalcatcher := make(chan os.Signal, 1)
	signal.Notify(signalcatcher, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-signalcatcher
}
