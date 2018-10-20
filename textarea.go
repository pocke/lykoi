package lykoi

import (
	"sync"
	"unicode/utf8"

	termbox "github.com/nsf/termbox-go"
)

type TextArea struct {
	renderer *Renderer
	mu       sync.Mutex
	buffer   *Buffer
}

func NewTextArea(b *Buffer, r *Renderer) (*TextArea, error) {
	t := &TextArea{
		renderer: r,
		buffer:   b,
	}
	go t.watch(b)
	err := t.Render()
	return t, err
}

func (c *TextArea) watch(b *Buffer) {
	ch := make(chan Buffer, 10)
	b.Subscribe(ch)
	for b := range ch {
		c.buffer = &b
		if err := c.Render(); err != nil {
			panic(err)
		}
	}
}

func (c *TextArea) Render() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	text := c.buffer.Text
	line := 0
	col := 0
	for len(text) > 0 {
		r, size := utf8.DecodeRune(text)
		text = text[size:]

		if r == '\n' {
			line += 1
			col = 0
			continue
		}
		err := c.renderer.SetCell(col, line, termbox.Cell{Ch: r})
		if err != nil {
			return err
		}

		col += 1
	}

	if err := c.renderer.Flush(); err != nil {
		return err
	}
	return nil
}

func (c *TextArea) UpdateRenderer(r *Renderer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.renderer = r
}

var _ Component = &TextArea{}
