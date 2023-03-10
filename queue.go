package main

type VoiceQueue []*TerentyVoice

var (
	TerentyVoiceQueue *VoiceQueue
)

func NewQueue() {
	TerentyVoiceQueue = &VoiceQueue{}
}

func (self *VoiceQueue) PushVoice(x *TerentyVoice) {
	*self = append(*self, x)
}

func (self *VoiceQueue) PopVoice() *TerentyVoice {
	h := *self
	var el *TerentyVoice
	l := len(h)

	if l == 0 {
		return nil
	}

	el, *self = h[0], h[1:l]

	// Or use this instead for a Stack
	// el, *self = h[l-1], h[0:l-1]
	return el
}
