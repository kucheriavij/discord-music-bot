package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/kucheriavij/discord-music-bot/src/handlers"
	"github.com/kucheriavij/discord-music-bot/src/queue"
	"github.com/kucheriavij/discord-music-bot/src/twitter"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token  string
	ApiUrl string
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("BOT_TOKEN")
	ApiUrl = os.Getenv("BOT_API_URL")

	queue.NewQueue()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)

	if err != nil {
		log.Fatal("Error creating Discord session", err)
	}

	dg.AddHandler(handlers.MessageCreate)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildMessages | discordgo.IntentsGuildVoiceStates)

	go twitter.Play(dg)

	err = dg.Open()

	if err != nil {
		log.Fatal("Error opening connection", err)
	}

	log.Print("Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
