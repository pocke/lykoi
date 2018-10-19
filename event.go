package lykoi

import (
	termbox "github.com/nsf/termbox-go"
)

func EventLoop() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				Exit(0)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
