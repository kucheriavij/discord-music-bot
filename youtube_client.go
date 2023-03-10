package main

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
)

func Search(query string, channel *discordgo.Channel) {
	service, err := youtube.NewService(context.Background(), option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))

	if err != nil {
		log.Printf("Error creating new YouTube client: %v", err)
		return
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id", "snippet"}).
		Q(query).
		MaxResults(1)
	response, err := call.Do()

	if err != nil {
		log.Printf("Response YouTube client: %v", err)
		return
	}

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			TerentyVoiceQueue.PushVoice(&TerentyVoice{
				Url:          fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId),
				VoiceChannel: channel,
			})
		}
	}
}
