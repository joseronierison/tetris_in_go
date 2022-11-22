package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pieceL boardPiece = boardPiece{
	pieceType: 'L',
	x:         24,
	y:         0,
	rotation:  0,
	width:     3,
	height:    3,
	isFalling: true,
}

var pieceT boardPiece = boardPiece{
	pieceType: 'T',
	x:         24,
	y:         0,
	rotation:  0,
	width:     3,
	height:    2,
	isFalling: true,
}

var pieceI boardPiece = boardPiece{
	pieceType: 'I',
	x:         24,
	y:         0,
	rotation:  0,
	width:     1,
	height:    3,
	isFalling: true,
}

var pieceS boardPiece = boardPiece{
	pieceType: 'S',
	x:         24,
	y:         0,
	rotation:  0,
	width:     2,
	height:    2,
	isFalling: true,
}

var pieceDot boardPiece = boardPiece{
	pieceType: '.',
	x:         24,
	y:         0,
	rotation:  0,
	width:     1,
	height:    1,
	isFalling: true,
}

func TestAllBoardFieldsAreEmpty(t *testing.T) {
	board := NewBoard(pieceL)
	assertBoardStateWithAnL(t, board.fields, pieceL.x, pieceL.y)
}

func TestFallingPieceInitialBoardPlacement(t *testing.T) {
	board := NewBoard(pieceL)
	fields := board.fields

	assertBoardStateWithAnL(t, fields, pieceL.x, pieceL.y)
}

func TestFalingPieceBehaviorWhenThereAreThreeTicks(t *testing.T) {
	board := NewBoard(pieceL)
	steps := 3
	for i := 0; i < steps; i++ {
		board.Tick()
	}

	piece := board.GetFallingPiece()

	assert.Equal(t, 24, piece.x)
	assert.Equal(t, steps, piece.y)
}

func TestThatFallingPieceGoesOnlyTilTheGround(t *testing.T) {
	board := NewBoard(pieceL)
	steps := 30
	for i := 0; i <= steps; i++ {
		board.Tick()
	}

	assert.True(t, board.fields[24][25])
	assert.True(t, board.fields[24][26])
	assert.True(t, board.fields[24][27])
	assert.True(t, board.fields[25][27])
	assert.True(t, board.fields[26][27])
}

func TestThatFallingPieceGoesTilTheGroundAfterBeenRotated(t *testing.T) {
	board := NewBoard(pieceL)
	steps := 30
	piece := board.GetFallingPiece()
	piece.rotate(&board)
	for i := 0; i <= steps; i++ {
		board.Tick()
	}

	assert.True(t, board.fields[24][27])
	assert.True(t, board.fields[25][27])
	assert.True(t, board.fields[26][27])
	assert.True(t, board.fields[26][25])
	assert.True(t, board.fields[26][26])
}

func TestThatTPiecePlacementInBoard(t *testing.T) {
	board := NewBoard(pieceT)

	assertBoardStateWithAT(t, board.fields, pieceT.x, pieceT.y)
}

func TestThatIPiecePlacementInBoard(t *testing.T) {
	board := NewBoard(pieceI)

	assertBoardStateWithAI(t, board.fields, pieceI.x, pieceI.y)
}

func TestThatSPiecePlacementInBoard(t *testing.T) {
	board := NewBoard(pieceS)

	assertBoardStateWithAS(t, board.fields, pieceS.x, pieceS.y)
}

func TestFallingPieceShouldStopOnceHitsAnotherPieceVertically(t *testing.T) {
	board := NewBoard(pieceS)
	board.nextPiece = pieceDot
	fallingPiece := board.GetFallingPiece()

	for i := 0; i < 49; i++ {
		for j := 5; j < 28; j++ {
			board.fields[i][j] = true
		}
	}

	board.Tick()
	board.Tick()
	board.Tick()
	board.Tick()

	for i := 0; i < 49; i++ {
		for j := 0; j < 27; j++ {
			if (i >= 24 && i <= 25 && j >= 3 && j <= 4) || j >= 5 || (i == fallingPiece.x && j == fallingPiece.y) {
				assert.True(t, board.fields[i][j], fmt.Sprintf("i: %v, j: %v", i, j))
			} else {
				assert.False(t, board.fields[i][j])
			}
		}
	}
}

func TestThatDotPiecePlacementInBoard(t *testing.T) {
	board := NewBoard(pieceDot)

	assertBoardStateWithADot(t, board.fields, pieceDot.x, pieceDot.y)
}

func assertBoardStateWithAnL(t *testing.T, fields [49][28]bool, x, y int) {
	for i := 0; i < len(fields); i++ {
		for j := 0; j < len(fields[i]); j++ {
			field := fields[i][j]
			if (i == x && j >= y && j <= y+2) || (j == y+2 && i >= x && i <= x+2) {
				// check the piece placement
				assert.True(t, field)
			} else {
				assert.False(t, field)
			}
		}
	}
}

func assertBoardStateWithAT(t *testing.T, fields [49][28]bool, x, y int) {
	for i := 0; i < len(fields); i++ {
		for j := 0; j < len(fields[i]); j++ {
			field := fields[i][j]
			if (j == y && i >= x && i <= x+2) || (i == x+1 && j >= y && j <= y+1) {
				// check the piece placement
				assert.True(t, field)
			} else {
				assert.False(t, field, fmt.Sprintf("i: %v, j: %v", i, j))
			}
		}
	}
}

func assertBoardStateWithAI(t *testing.T, fields [49][28]bool, x, y int) {
	for i := 0; i < len(fields); i++ {
		for j := 0; j < len(fields[i]); j++ {
			field := fields[i][j]
			if i == x && j >= y && j <= y+2 {
				// check the piece placement
				assert.True(t, field)
			} else {
				assert.False(t, field, fmt.Sprintf("i: %v, j: %v", i, j))
			}
		}
	}
}

func assertBoardStateWithAS(t *testing.T, fields [49][28]bool, x, y int) {
	for i := 0; i < len(fields); i++ {
		for j := 0; j < len(fields[i]); j++ {
			field := fields[i][j]
			if (i == x && j >= y && j <= y+1) || (i == x+1 && j >= y && j <= y+1) {
				// check the piece placement
				assert.True(t, field, fmt.Sprintf("i: %v, j: %v", i, j))
			} else {
				assert.False(t, field, fmt.Sprintf("i: %v, j: %v", i, j))
			}
		}
	}
}

func assertBoardStateWithADot(t *testing.T, fields [49][28]bool, x, y int) {
	for i := 0; i < len(fields); i++ {
		for j := 0; j < len(fields[i]); j++ {
			field := fields[i][j]
			if i == x && j == y {
				// check the piece placement
				assert.True(t, field, fmt.Sprintf("i: %v, j: %v", i, j))
			} else {
				assert.False(t, field, fmt.Sprintf("i: %v, j: %v", i, j))
			}
		}
	}
}

// func printBoardFields(fields [49][28]bool) {
// 	for j := 0; j < 28; j++ {
// 		var line string = ""
// 		for i := 0; i < 49; i++ {
// 			var field string
// 			if fields[i][j] {
// 				field = "1"
// 			} else {
// 				field = "0"
// 			}

// 			line = line + field
// 		}

// 		head := fmt.Sprintf("<- %v", j)

// 		println(line + head)
// 	}
// }
