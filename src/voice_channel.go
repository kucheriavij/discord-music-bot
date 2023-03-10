package src

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func GetVoiceChannel(s *discordgo.Session, guildId string, channelId string, m *discordgo.MessageCreate) *discordgo.Channel {
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
