package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kucheriavij/discord-music-bot/src/queue"
	"log"
	"regexp"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!play https://") {
		re := regexp.MustCompile("(https://www.youtube.com/watch\\?v=\\w+)")
		match := re.FindString(m.Content)

		if match == "" {
			log.Println("Alarma!!!!!")
			return
		}

		//_, err := s.ChannelMessageSend(m.ChannelID, match)
		//
		//if err != nil {
		//	log.Println("Error sending message to discord")
		//	return
		//}

		queue.TerentyQueue.Push(match)
	}
}
