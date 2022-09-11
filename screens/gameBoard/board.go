package gameboard

import (
	"os"
	"tetris/utils/graphic"

	"github.com/nsf/termbox-go"
)

func Init() {
	drawDashboard()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				termbox.Flush()
				os.Exit(1)
			}
		}
	}
}

func drawDashboard() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	width, heigth := graphic.GetScreenSize()

	temp := "The game will be here soon!"

	graphic.DrawBorders()

	graphic.WriteText((width-len(temp))/2, heigth/2, termbox.ColorLightCyan, termbox.ColorDefault, temp)

	termbox.Flush()
}
