package game

import (
	"tetris/game/graphic"

	"github.com/nsf/termbox-go"
)

func DrawL(x, y int) {
	graphic.DrawVerticalLine(x, y, 3, termbox.ColorBlue)
	graphic.DrawVerticalLine(x+1, y, 3, termbox.ColorBlue)
	graphic.DrawHorizontalLine(x+2, y+2, 2, termbox.ColorBlue)
}

func DrawI(x, y int) {
	graphic.DrawVerticalLine(x, y, 3, termbox.ColorBlue)
	graphic.DrawVerticalLine(x+1, y, 3, termbox.ColorBlue)
}

func DrawSquare(x, y int) {
	graphic.DrawHorizontalLine(x+2, y+2, 2, termbox.ColorBlue)
	graphic.DrawVerticalLine(x, y, 2, termbox.ColorBlue)
	graphic.DrawVerticalLine(x+1, y, 2, termbox.ColorBlue)
}

func DrawT(x, y int) {
	graphic.DrawHorizontalLine(x, y, 4, termbox.ColorBlue)
	graphic.DrawVerticalLine(x, y+1, 2, termbox.ColorBlue)
	graphic.DrawVerticalLine(x+1, y+1, 2, termbox.ColorBlue)
}
