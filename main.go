package lykoi

import (
	"io/ioutil"
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
	buf := State.Buffers[0]

	// test code
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return err
	}
	buf.Text = b

	w, h := termbox.Size()
	r := &Renderer{
		xOffset: 0, yOffset: 0,
		width: w, height: h,
	}
	_, err = NewTextArea(buf, r)
	if err != nil {
		return err
	}

	return nil
}

func Exit(status int) {
	termbox.Close()
	os.Exit(status)
}
