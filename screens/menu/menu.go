package menu

import (
	"os"
	"tetris/screens"
	"tetris/utils/graphic"

	"github.com/nsf/termbox-go"
)

var menuState = 0

func Init() {
	drawMenu()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				os.Exit(1)
			case termbox.KeyArrowUp:
				if menuState > 0 {
					menuState = menuState - 1
					drawMenu()
				}
			case termbox.KeyArrowDown:
				if menuState < 1 {
					menuState = menuState + 1
					drawMenu()
				}
			case termbox.KeyEnter:
				switch menuState {
				case 0:
					screens.SS.GoToGame()
					break mainloop
				case 1:
					termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
					termbox.Flush()
					os.Exit(1)
				}
			}
		}
	}
}

func drawMenu() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	width, heigth := graphic.GetScreenSize()

	const menuWidth = 30
	const menuHeigth = 12
	x := (width - menuWidth) / 2
	y := (heigth - menuHeigth) / 2

	title := "This is a Tetris written in GoLang!"
	copyright := "Made by Roni Silva"
	github := "https://github.com/joseronierison"

	graphic.DrawBorders()

	graphic.WriteText((width-len(title))/2, y-3, termbox.ColorLightCyan, termbox.ColorDefault, title)
	graphic.WriteText((width-len(copyright))/2, y+15, termbox.ColorLightGreen, termbox.ColorDefault, copyright)
	graphic.WriteText((width-len(github))/2, y+16, termbox.ColorLightBlue, termbox.ColorDefault, github)

	newGameText := "NEW GAME"
	exitGameText := "EXIT"
	selectChar := "▶"

	if menuState == 0 {
		graphic.WriteText((width-len(newGameText))/2, y+4, termbox.ColorRed, termbox.ColorDefault, newGameText)
		graphic.WriteText((width-len(newGameText))/2-2, y+4, termbox.ColorRed, termbox.ColorDefault, selectChar)
		graphic.WriteText((width-len(exitGameText))/2, y+7, termbox.ColorWhite, termbox.ColorDefault, exitGameText)
	}

	if menuState == 1 {
		graphic.WriteText((width-len(newGameText))/2, y+4, termbox.ColorWhite, termbox.ColorDefault, newGameText)
		graphic.WriteText((width-len(exitGameText))/2, y+7, termbox.ColorRed, termbox.ColorDefault, exitGameText)
		graphic.WriteText((width-len(exitGameText))/2-2, y+7, termbox.ColorRed, termbox.ColorDefault, selectChar)
	}

	graphic.DrawRectangle(x, y, menuWidth, menuHeigth, termbox.ColorDarkGray)
	termbox.Flush()
}
