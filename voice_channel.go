package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"regexp"
)

func GetVoiceChannel(s *discordgo.Session, guildId string, channelId string, m *discordgo.MessageCreate) *discordgo.Channel {
	if m.Author == nil || m.Author.Bot || !isPlayCommand(m.Content) {
		return nil
	}

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

func isPlayCommand(command string) bool {
	re := regexp.MustCompile("^!(search|play|random|stop)\\b")
	match := re.FindString(command)

	return match != ""
}
