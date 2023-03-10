package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var (
	Token string
)

func init() {
	executable, err := os.Executable()

	if err != nil {
		log.Fatal("Cant get current executable path")
	}

	dir, err := filepath.Abs(executable)

	if err != nil {
		log.Fatal(err)
	}

	dir = filepath.Dir(dir)

	err = os.MkdirAll(fmt.Sprintf("%s/var/log", dir), 0755)

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/var/log/terenty.log", dir), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	err = godotenv.Load(fmt.Sprintf("%s/.env", dir))

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
