package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"regexp"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author == nil {
		log.Println("Отправил никто")
		return
	}

	channel := GetVoiceChannel(s, m.GuildID, m.ChannelID, m)

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

		TerentyVoiceQueue.PushVoice(&TerentyVoice{
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

		Search(m.Content, channel)
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
