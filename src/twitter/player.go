package twitter // Package twitter "Щебетарь"

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/jonas747/dca"
	"github.com/kkdai/youtube/v2"
	"github.com/kucheriavij/discord-music-bot/src/queue"
	"log"
	"os"
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

		log.Println("Начинаю грузить апельсины бочками.")
		fileName, err := downloadFile(fmt.Sprint(url))
		log.Println("Найдено 0 вирусов. Я так и думал.")

		if err != nil {
			continue
		}

		encodeSession, err := dca.EncodeFile(fileName, dca.StdEncodeOptions)

		closeSession := func() {
			encodeSession.Cleanup()
		}

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
				closeSession()
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

				closeSession()
				break
			}
		}

		continue
	}
}

func downloadFile(url string) (string, error) {
	c := youtube.Client{}

	video, err := c.GetVideo(fmt.Sprint(url))

	if err != nil {
		log.Println("Get video error", err)
		return "", fmt.Errorf("get video error")
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
		return "", fmt.Errorf("format not found")
	}

	closure, length, err := c.GetStream(video, format)

	if err != nil {
		log.Println("Can not get stream", err)
		return "", fmt.Errorf("can not get stream")
	}

	var counter int64 = 0
	buf := make([]byte, 1024)
	fileName := fmt.Sprintf("%s.opus", uuid.New().String())
	out, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)

	for {
		if counter >= length {
			break
		}

		read, err := closure.Read(buf)

		if err != nil {
			log.Println("Can not read video from youtube: ", err)
			err := out.Close()
			if err != nil {
				log.Printf("Can not close file: %s\n\r", fileName)
			} else {
				deleteFile(fileName)
			}
			return "", fmt.Errorf("can not read video from youtube")
		}

		_, err = out.Write(buf)

		if err != nil {
			log.Println("Can not write file: ", err)
			err := out.Close()
			if err != nil {
				log.Printf("Can not close file: %s\n\r", fileName)
			} else {
				deleteFile(fileName)
			}
			return "", fmt.Errorf("can not write file")
		}

		counter += int64(read)
	}

	err = out.Close()
	if err != nil {
		log.Println("Can not close file")
	}

	return fileName, nil
}

func deleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Printf("Can not delete file: %s\n\r", fileName)
	}
}
