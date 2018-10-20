package lykoi

import (
	"github.com/pkg/errors"
)

type StateT struct {
	currentBufNum int
	Buffers       []*Buffer
}

type Buffer struct {
	subscriptions []chan<- Buffer
	Text          []byte
}

func (b *Buffer) Subscribe(ch chan<- Buffer) {
	b.subscriptions = append(b.subscriptions, ch)
}

func (b *Buffer) Unsubscribe(ch chan<- Buffer) error {
	for idx, c := range b.subscriptions {
		if c == ch {
			b.subscriptions = append(b.subscriptions[:idx], b.subscriptions[idx+1:]...)
			return nil
		}
	}
	return errors.New("Unknown channel. Unsubscribe is failed")
}

func (b *Buffer) Update(f func(*Buffer)) {
	f(b)
	for _, c := range b.subscriptions {
		c <- *b
	}
}

var State StateT

func init() {
	State = StateT{
		Buffers: []*Buffer{
			{},
		},
	}
}

func (s *StateT) updateCurrentBuffer(f func(*Buffer)) {
	s.Buffers[s.currentBufNum].Update(f)
}
