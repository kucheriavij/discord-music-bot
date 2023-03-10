package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kucheriavij/discord-music-bot/src"
	"github.com/kucheriavij/discord-music-bot/src/queue"
	"github.com/kucheriavij/discord-music-bot/src/structures"
	"github.com/kucheriavij/discord-music-bot/src/youtube"
	"log"
	"regexp"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author == nil {
		log.Println("Отправил никто")
		return
	}

	channel := src.GetVoiceChannel(s, m.GuildID, m.ChannelID, m)

	if channel == nil {
		log.Println("Пользователь еще не сидит ☭")
		return
	}

	if strings.HasPrefix(m.Content, "!play https://") {
		re := regexp.MustCompile("((?:https?:)?\\/\\/)?((?:www|m)\\.)?((?:youtube(-nocookie)?\\.com|youtu.be))(\\/(?:[\\w\\-]+\\?v=|embed\\/|v\\/)?)([\\w\\-]+)(\\S+)?")
		match := re.FindString(m.Content)

		if match == "" {
			log.Println(m.Content)
			return
		}

		queue.TerentyVoiceQueue.PushVoice(&structures.TerentyVoice{
			Url:          match,
			VoiceChannel: channel,
		})
	}

	if strings.HasPrefix(m.Content, "!search ") {
		re := regexp.MustCompile("^!search\\s+(.{3,})$")
		match := re.FindString(m.Content)

		if match == "" {
			log.Println(m.Content)
			return
		}

		youtube.Search(m.Content, channel)
		log.Printf("query string: '%s'; query length: '%d'", m.Content, len(m.Content))
	}

	if strings.HasPrefix(m.Content, "!stop") {
		for _, item := range s.VoiceConnections {
			err := item.Disconnect()
			if err != nil {
				return
			}
		}
	}
}
