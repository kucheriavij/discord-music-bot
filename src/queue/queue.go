package queue

type Queue []interface{}

var TerentyQueue *Queue

func NewQueue() {
	TerentyQueue = &Queue{}
}

func (self *Queue) Push(x interface{}) {
	*self = append(*self, x)
}

func (self *Queue) Pop() interface{} {
	h := *self
	var el interface{}
	l := len(h)

	if l == 0 {
		return nil
	}

	el, *self = h[0], h[1:l]

	// Or use this instead for a Stack
	// el, *self = h[l-1], h[0:l-1]
	return el
}
