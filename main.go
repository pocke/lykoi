package lykoi

import (
	"os"

	termbox "github.com/nsf/termbox-go"
)

func Init() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	go EventLoop()

	return nil
}

func Exit(status int) {
	termbox.Close()
	os.Exit(status)
}
