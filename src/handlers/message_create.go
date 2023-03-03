package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kucheriavij/discord-music-bot/src/queue"
	"github.com/kucheriavij/discord-music-bot/src/structures"
	"log"
	"regexp"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!play https://") {
		re := regexp.MustCompile("(https://www.youtube.com/watch\\?v=\\w+)")
		match := re.FindString(m.Content)

		if match == "" {
			log.Println(m.Content)
			return
		}

		if m.Author == nil {
			log.Println("Отправил никто")
			return
		}

		channel := getVoiceChannel(s, m.GuildID, m.ChannelID, m)

		if channel == nil {
			log.Println("Пользователь еще не сидит ☭")
			return
		}

		queue.TerentyQueue.Push(&structures.TerentyVoice{
			Url:          match,
			VoiceChannel: channel,
		})
	}
}

func getVoiceChannel(s *discordgo.Session, guildId string, channelId string, m *discordgo.MessageCreate) *discordgo.Channel {
	channels, err := s.GuildChannels(guildId)

	if err != nil {
		log.Println("Не смог получить каналы")
		return nil
	}

	for _, item := range channels {
		if item.Type == discordgo.ChannelTypeGuildVoice && item.ID == channelId {
			return item
		}
	}

	_, err = s.ChannelMessageSendReply(channelId, "Иди в голосовой канал", m.Reference())
	if err != nil {
		log.Println("Иди в голосовой канал")
	}

	return nil
}
