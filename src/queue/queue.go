package queue

import "github.com/kucheriavij/discord-music-bot/src/structures"

type Queue []*structures.TerentyVoice

var TerentyQueue *Queue

func NewQueue() {
	TerentyQueue = &Queue{}
}

func (self *Queue) Push(x *structures.TerentyVoice) {
	*self = append(*self, x)
}

func (self *Queue) Pop() *structures.TerentyVoice {
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
