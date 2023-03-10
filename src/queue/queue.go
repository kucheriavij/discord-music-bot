package queue

import "github.com/kucheriavij/discord-music-bot/src/structures"

type VoiceQueue []*structures.TerentyVoice

var (
	TerentyVoiceQueue *VoiceQueue
)

func NewQueue() {
	TerentyVoiceQueue = &VoiceQueue{}
}

func (self *VoiceQueue) PushVoice(x *structures.TerentyVoice) {
	*self = append(*self, x)
}

func (self *VoiceQueue) PopVoice() *structures.TerentyVoice {
	h := *self
	var el *structures.TerentyVoice
	l := len(h)

	if l == 0 {
		return nil
	}

	el, *self = h[0], h[1:l]

	// Or use this instead for a Stack
	// el, *self = h[l-1], h[0:l-1]
	return el
}
