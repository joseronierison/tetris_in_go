package graphic

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func WriteText(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func GetScreenSize() (width, height int) {
	return 75, 30
}

func DrawExternalBoarders() {
	const coldef = termbox.ColorDefault
	w, h := GetScreenSize()

	DrawRectangle(0, 0, w, h, termbox.ColorWhite)
}

func DrawRectangle(x, y, width, height int, color termbox.Attribute) {
	drawContinuously(x, y, width, 1, termbox.Cell{Ch: ' ', Bg: color})          //top
	drawContinuously(x+width-1, y, 1, height, termbox.Cell{Ch: ' ', Bg: color}) //right
	drawContinuously(x, y+height-1, width, 1, termbox.Cell{Ch: ' ', Bg: color}) //bottom
	drawContinuously(x, y, 1, height, termbox.Cell{Ch: ' ', Bg: color})         //left
}

func DrawVerticalLine(x, y, size int, color termbox.Attribute) {
	drawContinuously(x, y, 1, size, termbox.Cell{Ch: ' ', Bg: color})
}

func DrawHorizontalLine(x, y, size int, color termbox.Attribute) {
	drawContinuously(x, y, size, 1, termbox.Cell{Ch: ' ', Bg: color})
}

func drawContinuously(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}
