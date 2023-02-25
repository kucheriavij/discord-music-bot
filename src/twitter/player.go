package twitter // Package twitter "Щебетарь"

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"github.com/kkdai/youtube/v2"
	"github.com/kucheriavij/discord-music-bot/src/queue"
	"log"
	"time"
)

func Play(s *discordgo.Session) {
	for {
		time.Sleep(500 * time.Millisecond)

		url := queue.TerentyQueue.Pop()

		if url == nil {
			continue
		}

		voice, err := s.ChannelVoiceJoin("947897321583149127", "947897322048741448", false, false)

		if err != nil {
			log.Println("Voice connection error: ", err)
			continue
		}

		for !voice.Ready {
			time.Sleep(10 * time.Millisecond)
		}

		log.Println("Connected. Ванька, думай над своим поведением.")

		encodeSession, err := dca.EncodeFile("test.mp3", dca.StdEncodeOptions)
		defer encodeSession.Cleanup()

		if err != nil {
			log.Println("Encoding file error: ", err)
		}

		for {
			frame, err := encodeSession.OpusFrame()

			if err != nil {
				voice.Close()
				err := voice.Disconnect()

				if err != nil {
					log.Println("Voice disconnected")
				}

				log.Println("End of file")
				break
			}

			select {
			case voice.OpusSend <- frame:
			case <-time.After(time.Second):
				voice.Close()
				err := voice.Disconnect()

				if err != nil {
					log.Println("Voice disconnected")
				}

				break
			}
		}

		continue

		c := youtube.Client{}

		video, err := c.GetVideo(fmt.Sprint(url))

		if err != nil {
			log.Println("Get video error", err)
			continue
		}

		log.Println(video.Formats)

		///*closer*/ _, length, err := c.GetStream(video, "")
		//
		//log.Println("Length", length)
		//
		//if err != nil {
		//	log.Println("Alarma!", err)
		//	continue
		//}
	}
}
