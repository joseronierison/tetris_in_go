package game

import (
	"math/rand"
	"tetris/game/graphic"
	"time"

	"github.com/nsf/termbox-go"
)

var letters = []rune{'r', 'o', 'n', 'i'}

func InitGameBoard(ss *ScreenState) {
	termbox.Init()
	termbox.SetInputMode(termbox.InputEsc)
	go someThreadTest()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				ss.Exit()
				return
			}
		}
	}
}

func drawDashboard() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	graphic.DrawExternalBoarders()

	termbox.Flush()
}

func drawGameInfoMenu() {
	graphic.DrawVerticalLine(50, 0, 30, termbox.ColorWhite)
}

func drawScore() {
	graphic.WriteText(55, 3, termbox.ColorWhite, termbox.ColorDefault, "Score")
	graphic.WriteText(65, 3, termbox.ColorWhite, termbox.ColorDefault, "000")
}

func someThreadTest() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	drawDashboard()
	drawGameInfoMenu()
	drawScore()
	DrawL(60, 5)

	for {
		for x := 10; x < 20; x++ {
			for y := 10; y < 20; y++ {
				termbox.SetChar(x, y, letters[rand.Intn(len(letters))])
				termbox.SetFg(x, y, termbox.Attribute(rand.Intn(9)))
			}
		}
		termbox.Flush()
		<-ticker.C
	}
}
