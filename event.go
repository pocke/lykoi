package lykoi

import (
	"unicode/utf8"

	termbox "github.com/nsf/termbox-go"
)

func EventLoop() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				Exit(0)
			default:
				if ev.Ch != 0 {
					State.updateCurrentBuffer(func(b *Buffer) {
						ch := runeToBytes(ev.Ch)

						b.Text = append(b.Text, ch...)
					})
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func runeToBytes(r rune) []byte {
	var buf [utf8.UTFMax]byte
	size := utf8.EncodeRune(buf[:], r)
	return buf[:size]
}
