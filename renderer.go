package lykoi

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/pkg/errors"
)

type Renderer struct {
	yOffset int
	xOffset int
	height  int
	width   int
}

func (r *Renderer) SetCell(x, y int, cell termbox.Cell) error {
	if x < r.xOffset || r.width < x {
		return errors.Errorf("x should be %d < x < %d, but x == %d", r.xOffset, r.width, x)
	}
	if y < r.yOffset || r.height < y {
		return errors.Errorf("y should be %d < y < %d, but y == %d", r.yOffset, r.height, y)
	}

	termbox.SetCell(x+r.xOffset, y+r.yOffset, cell.Ch, cell.Fg, cell.Bg)
	return nil
}

func (r *Renderer) Width() int {
	return r.width - r.xOffset
}

func (r *Renderer) Height() int {
	return r.height - r.yOffset
}

func (r *Renderer) Fill(x1, y1, x2, y2 int, cell termbox.Cell) error {
	if x1 > x2 {
		return errors.Errorf("Fill: x1(%d) > x2(%d)", x1, x2)
	}
	if y1 > y2 {
		return errors.Errorf("Fill: y1(%d) > y2(%d)", y1, y2)
	}

	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			if err := r.SetCell(x, y, cell); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Renderer) Flush() error {
	return termbox.Flush()
}
