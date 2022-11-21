package screens

import (
	"tetris/game/graphic"

	"github.com/nsf/termbox-go"
)

var menuState = 0

func InitMenu(ss *ScreenState) {
	drawMenu()
	ch := make(chan bool)
	go drawAnimation(ch)

menuloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				ss.Exit()
				break menuloop
			case termbox.KeyArrowUp:
				if menuState > 0 {
					menuState = menuState - 1
				}
			case termbox.KeyArrowDown:
				if menuState < 1 {
					menuState = menuState + 1
				}
			case termbox.KeyEnter:
				switch menuState {
				case 0:
					close(ch)
					ss.GoToGame()
					break menuloop
				case 1:
					ss.Exit()
					break menuloop
				}
			}
		}

		drawMenu()
	}
}

func drawMenu() {
	width, heigth := graphic.GetScreenSize()

	const menuWidth = 30
	const menuHeigth = 12
	x := (width - menuWidth) / 2
	y := (heigth - menuHeigth) / 2

	title := "This is a Tetris written in GoLang!"
	copyright := "Made by Roni Silva"
	github := "https://github.com/joseronierison"

	graphic.DrawExternalBoarders()

	graphic.WriteText((width-len(title))/2, y-3, termbox.ColorLightCyan, termbox.ColorDefault, title)
	graphic.WriteText((width-len(copyright))/2, y+15, termbox.ColorLightGreen, termbox.ColorDefault, copyright)
	graphic.WriteText((width-len(github))/2, y+16, termbox.ColorLightBlue, termbox.ColorDefault, github)

	newGameText := "NEW GAME"
	exitGameText := "EXIT"
	selectChar := "▶"

	if menuState == 0 {
		graphic.WriteText((width-len(newGameText))/2, y+4, termbox.ColorRed, termbox.ColorDefault, newGameText)
		graphic.WriteText((width-len(exitGameText))/2, y+7, termbox.ColorWhite, termbox.ColorDefault, exitGameText)
		graphic.WriteText((width-len(newGameText))/2-2, y+4, termbox.ColorRed, termbox.ColorDefault, selectChar)
		graphic.WriteText((width-len(exitGameText))/2-2, y+7, termbox.ColorDefault, termbox.ColorDefault, " ")
	}

	if menuState == 1 {
		graphic.WriteText((width-len(newGameText))/2, y+4, termbox.ColorWhite, termbox.ColorDefault, newGameText)
		graphic.WriteText((width-len(exitGameText))/2, y+7, termbox.ColorRed, termbox.ColorDefault, exitGameText)
		graphic.WriteText((width-len(exitGameText))/2-2, y+7, termbox.ColorRed, termbox.ColorDefault, selectChar)
		graphic.WriteText((width-len(newGameText))/2-2, y+4, termbox.ColorDefault, termbox.ColorDefault, " ")
	}

	graphic.DrawRectangle(x, y, menuWidth, menuHeigth, termbox.ColorDarkGray)
	termbox.Flush()
}

func drawAnimation(ch chan bool) {
	var doa = graphic.DrawableAtom{Char: ' ', Fg: termbox.ColorDefault, Bg: termbox.ColorGreen}
	var eao = graphic.DrawableAtom{Char: ' ', Fg: termbox.ColorDefault, Bg: termbox.ColorDefault}

	var t1Object = graphic.DrawableObject{Atoms: [][]graphic.DrawableAtom{{doa, doa, doa, doa, doa}, {eao, eao, doa}, {eao, eao, doa}, {eao, eao, doa}, {eao, eao, doa}}}
	var eObject = graphic.DrawableObject{Atoms: [][]graphic.DrawableAtom{{doa, doa, doa, doa}, {doa}, {doa, doa, doa, eao}, {doa}, {doa, doa, doa, doa}}}
	var t2Object = graphic.DrawableObject{Atoms: [][]graphic.DrawableAtom{{doa, doa, doa, doa, doa}, {eao, eao, doa}, {eao, eao, doa}, {eao, eao, doa}, {eao, eao, doa}}}
	var rObject = graphic.DrawableObject{Atoms: [][]graphic.DrawableAtom{{doa, doa, doa, doa}, {doa, eao, eao, doa}, {doa, doa, doa, doa}, {doa, eao, doa, eao}, {doa, eao, eao, doa}}}
	var iObject = graphic.DrawableObject{Atoms: [][]graphic.DrawableAtom{{doa, doa, doa}, {eao, doa}, {eao, doa}, {eao, doa}, {doa, doa, doa}}}
	var sObject = graphic.DrawableObject{Atoms: [][]graphic.DrawableAtom{{doa, doa, doa, doa}, {doa}, {doa, doa, doa, doa}, {eao, eao, eao, doa}, {doa, doa, doa, doa}}}

	var rightOrigin = graphic.Coordinates{X: 65, Y: 2}
	var leftOrigin = graphic.Coordinates{X: 59, Y: 2}
	var baseSteps = 19

	sObject.MoveObjectVertically(rightOrigin, baseSteps, ch)
	iObject.MoveObjectVertically(leftOrigin, baseSteps, ch)
	rObject.MoveObjectVertically(rightOrigin, baseSteps-6, ch)
	t2Object.MoveObjectVertically(leftOrigin, baseSteps-6, ch)
	eObject.MoveObjectVertically(rightOrigin, baseSteps-12, ch)
	t1Object.MoveObjectVertically(leftOrigin, baseSteps-12, ch)
}
