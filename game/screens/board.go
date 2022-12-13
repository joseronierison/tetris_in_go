package screens

import (
	"strconv"
	"tetris/game/core"
	"tetris/game/graphic"
	"time"

	"github.com/nsf/termbox-go"
)

var tetrisBoard = core.NewBoard(core.GenerateRandomFallingPiece())

func InitGameBoard(ss *ScreenState) {
	termbox.Init()
	termbox.SetInputMode(termbox.InputEsc)
	go drawBoardFrames()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				ss.Exit()
				return
			case termbox.KeyArrowLeft:
				tetrisBoard.GetFallingPiece().MoveLeft(&tetrisBoard)
			case termbox.KeyArrowRight:
				tetrisBoard.GetFallingPiece().MoveRight(&tetrisBoard)
			case termbox.KeySpace:
				tetrisBoard.GetFallingPiece().Rotate(&tetrisBoard)
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
	graphic.WriteText(67, 3, termbox.ColorWhite, termbox.ColorDefault, strconv.Itoa(tetrisBoard.GetScore()))
}

func drawNextPiece() {
	coordinates := graphic.Coordinates{X: 67, Y: 5}
	nextPiece := tetrisBoard.GetNexPiece()
	graphic.EmptyObject.DrawObject(coordinates)
	graphic.WriteText(55, 5, termbox.ColorWhite, termbox.ColorDefault, "Next Piece:")
	switch nextPiece.GetPieceType() {
	case 'L':
		graphic.LObject.DrawObject(coordinates)
	case 'I':
		graphic.IObject.DrawObject(coordinates)
	case 'S':
		graphic.SObject.DrawObject(coordinates)
	case 'T':
		graphic.TObject.DrawObject(coordinates)
	case '.':
		graphic.DotObject.DrawObject(coordinates)
	}
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

func drawBoardFrames() {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

	drawExternalDashboard()
	drawGameInfoMenu()
	drawMenuInstructions()
	for {

		drawGameBoard()

		drawScore()
		drawNextPiece()
		tetrisBoard.Tick()

		if tetrisBoard.IsOver() {
			graphic.DrawRectangle(20, 10, 14, 3, termbox.ColorLightRed)
			graphic.WriteText(22, 11, termbox.ColorRed, termbox.ColorDefault, "Game Over!")
		}

		termbox.Flush()
		<-ticker.C
	}
}
