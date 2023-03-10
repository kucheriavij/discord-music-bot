package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token string
)

func init() {
	err := os.MkdirAll("var/log", 0755)

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("var/log/terenty.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("BOT_TOKEN")

	NewQueue()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)

	if err != nil {
		log.Fatal("Error creating Discord session", err)
	}

	dg.AddHandler(MessageCreate)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildMessages | discordgo.IntentsGuildVoiceStates | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences)

	go Play(dg)

	err = dg.Open()

	if err != nil {
		log.Fatal("Error opening connection", err)
	}

	log.Println("Бот стартанул")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = dg.Close()
	if err != nil {
		log.Println("Не смог закрыть")
	}
}
