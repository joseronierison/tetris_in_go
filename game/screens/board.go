package screens

import (
	"strconv"
	"tetris/game/core"
	"tetris/game/graphic"
	"time"

	"github.com/nsf/termbox-go"
)

var tetrisBoard = core.NewBoard(core.GenerateRandomFallingPiece())
var score int = 0

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

func drawExternalDashboard() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	graphic.DrawExternalBoarders()

	termbox.Flush()
}

func drawGameInfoMenu() {
	graphic.DrawVerticalLine(50, 0, 30, termbox.ColorWhite)
}

func drawScore() {
	graphic.WriteText(55, 3, termbox.ColorWhite, termbox.ColorDefault, "Score:")
	graphic.WriteText(65, 3, termbox.ColorWhite, termbox.ColorDefault, strconv.Itoa(score))
}

func drawMenuInstructions() {
	baseY := 23
	graphic.WriteText(52, baseY, termbox.ColorWhite, termbox.ColorDefault, "← or → to move")
	graphic.WriteText(52, baseY+1, termbox.ColorWhite, termbox.ColorDefault, "↓ to acelerate")
	graphic.WriteText(52, baseY+4, termbox.ColorWhite, termbox.ColorDefault, "Espace to rotate")
	graphic.WriteText(52, baseY+2, termbox.ColorWhite, termbox.ColorDefault, "p to pause")
	graphic.WriteText(52, baseY+3, termbox.ColorWhite, termbox.ColorDefault, "r to resume")
	graphic.WriteText(52, baseY+5, termbox.ColorWhite, termbox.ColorDefault, "ESC to quit")
}

func drawGameBoard() {
	for lineIndex, column := range tetrisBoard.GetFields() {
		x := lineIndex + 1
		for columnIndex, field := range column {
			y := columnIndex + 1

			if field {
				termbox.SetBg(x, y, termbox.ColorLightGray)
			} else {
				termbox.SetBg(x, y, termbox.ColorBlack)
			}
		}
	}
}

func someThreadTest() {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

	drawExternalDashboard()
	drawGameInfoMenu()
	drawMenuInstructions()
	//DrawL(60, 5)

	for {
		drawGameBoard()

		drawScore()

		tetrisBoard.Tick()

		termbox.Flush()
		<-ticker.C
	}
}
