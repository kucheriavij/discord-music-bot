package twitter // Package twitter - "Щебетарь"

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/kkdai/youtube/v2"
	"github.com/kucheriavij/discord-music-bot/src/queue"
	"github.com/kucheriavij/discord-music-bot/src/structures"
	"io"
	"log"
	"time"
)

func Play(s *discordgo.Session) {
	for {
		time.Sleep(500 * time.Millisecond)

		terenty := queue.TerentyVoiceQueue.PopVoice()

		if terenty == nil {
			continue
		}

		playTweet(s, terenty)
	}
}

func playTweet(s *discordgo.Session, terenty *structures.TerentyVoice) {
	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = terenty.VoiceChannel.Bitrate / 1000
	options.Application = "lowdelay"

	voice, err := s.ChannelVoiceJoin(terenty.VoiceChannel.GuildID, terenty.VoiceChannel.ID, false, false)

	if err != nil {
		log.Println("Voice connection error: ", err)
		return
	}

	defer voice.Close()
	defer func(voice *discordgo.VoiceConnection) {
		err := voice.Disconnect()
		if err != nil {
			log.Println(err)
		}
	}(voice)

	for !voice.Ready {
		time.Sleep(10 * time.Millisecond)
	}

	c := youtube.Client{}

	video, err := c.GetVideo(fmt.Sprint(terenty.Url))

	if err != nil {
		log.Println("Get video error", err)
		return
	}

	var format *youtube.Format
	formats := video.Formats

	for _, v := range formats {
		if v.MimeType == "audio/webm; codecs=\"opus\"" {
			format = &v
			break
		}
	}

	if format == nil {
		log.Println("Нет поддерживаемого формата")
		return
	}

	url, err := c.GetStreamURL(video, format)
	if err != nil {
		log.Println("Can not get stream", err)
		return
	}

	encodingSession, err := dca.EncodeFile(url, options)
	if err != nil {
		log.Println(err)
	}
	defer encodingSession.Cleanup()

	done := make(chan error)
	dca.NewStream(encodingSession, voice, done)
	err = <-done
	if err != nil && err != io.EOF {
		_, err = s.ChannelMessageSend(terenty.VoiceChannel.ID, "Мммм, хуета")
		log.Println("Мммм, хуета")
	}
}
