package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFallingPieceInitialDefinition(t *testing.T) {
	board := NewBoard(pieceL)
	piece := board.GetFallingPiece()

	assert.True(t, piece.isFalling)
	assert.Equal(t, piece.x, 24)
	assert.Equal(t, piece.y, 0)
}

func TestPieceRotationsInATFallingPiece(t *testing.T) {
	board := NewBoard(pieceT)
	x := pieceT.x
	y := pieceT.y
	assert.Equal(t, 0, pieceT.rotation)
	assert.Equal(t, 3, pieceT.width)
	assert.Equal(t, 2, pieceT.height)
	assertBoardStateWithAT(t, board.fields, pieceT.x, pieceT.y)

	pieceT.Rotate(&board)
	assert.Equal(t, 90, pieceT.rotation)
	assert.Equal(t, 2, pieceT.width)
	assert.Equal(t, 3, pieceT.height)
	assert.True(t, board.fields[x][y])
	assert.True(t, board.fields[x][y+1])
	assert.True(t, board.fields[x+1][y+1])
	assert.True(t, board.fields[x][y+2])

	pieceT.Rotate(&board)
	assert.Equal(t, 180, pieceT.rotation)
	assert.Equal(t, 3, pieceT.width)
	assert.Equal(t, 2, pieceT.height)
	assert.True(t, board.fields[x][y+1])
	assert.True(t, board.fields[x+1][y+1])
	assert.True(t, board.fields[x+1][y])
	assert.True(t, board.fields[x+2][y+1])

	pieceT.Rotate(&board)
	assert.Equal(t, 270, pieceT.rotation)
	assert.Equal(t, 2, pieceT.width)
	assert.Equal(t, 3, pieceT.height)
	assert.True(t, board.fields[x+1][y])
	assert.True(t, board.fields[x+1][y+1])
	assert.True(t, board.fields[x][y+1])
	assert.True(t, board.fields[x+1][y+2])

	pieceT.Rotate(&board)
	assert.Equal(t, 0, pieceT.rotation)
	assert.Equal(t, 3, pieceT.width)
	assert.Equal(t, 2, pieceT.height)
	assertBoardStateWithAT(t, board.fields, pieceT.x, pieceT.y)
}

func TestPieceRotationsInAnLFallingPiece(t *testing.T) {
	board := NewBoard(pieceL)
	x := pieceL.x
	y := pieceL.y
	assert.Equal(t, 0, pieceL.rotation)
	assert.Equal(t, 3, pieceL.width)
	assert.Equal(t, 3, pieceL.height)
	assertBoardStateWithAnL(t, board.fields, pieceL.x, pieceL.y)

	pieceL.Rotate(&board)
	assert.Equal(t, 90, pieceL.rotation)
	assert.Equal(t, 3, pieceL.width)
	assert.Equal(t, 3, pieceL.height)
	assert.True(t, board.fields[x][y+2])
	assert.True(t, board.fields[x+1][y+2])
	assert.True(t, board.fields[x+2][y+2])
	assert.True(t, board.fields[x+2][y+1])
	assert.True(t, board.fields[x+2][y])

	pieceL.Rotate(&board)
	assert.Equal(t, 180, pieceL.rotation)
	assert.Equal(t, 3, pieceL.width)
	assert.Equal(t, 3, pieceL.height)
	assert.True(t, board.fields[x][y])
	assert.True(t, board.fields[x+1][y])
	assert.True(t, board.fields[x+2][y])
	assert.True(t, board.fields[x+2][y+1])
	assert.True(t, board.fields[x+2][y+2])

	pieceL.Rotate(&board)
	assert.Equal(t, 270, pieceL.rotation)
	assert.Equal(t, 3, pieceL.width)
	assert.Equal(t, 3, pieceL.height)
	assert.True(t, board.fields[x][y])
	assert.True(t, board.fields[x+1][y])
	assert.True(t, board.fields[x+2][y])
	assert.True(t, board.fields[x][y+1])
	assert.True(t, board.fields[x][y+2])

	pieceL.Rotate(&board)
	assert.Equal(t, 0, pieceL.rotation)
	assert.Equal(t, 3, pieceL.width)
	assert.Equal(t, 3, pieceL.height)
	assertBoardStateWithAnL(t, board.fields, pieceL.x, pieceL.y)
}

func TestPieceRotationsInAIFallingPiece(t *testing.T) {
	board := NewBoard(pieceI)
	x := pieceI.x
	y := pieceI.y
	assert.Equal(t, 0, pieceI.rotation)
	assert.Equal(t, 1, pieceI.width)
	assert.Equal(t, 3, pieceI.height)
	assertBoardStateWithAI(t, board.fields, pieceI.x, pieceI.y)

	pieceI.Rotate(&board)
	assert.Equal(t, 90, pieceI.rotation)
	assert.Equal(t, 3, pieceI.width)
	assert.Equal(t, 1, pieceI.height)
	assert.True(t, board.fields[x][y])
	assert.True(t, board.fields[x+1][y])
	assert.True(t, board.fields[x+2][y])

	pieceI.Rotate(&board)
	assert.Equal(t, 180, pieceI.rotation)
	assert.Equal(t, 1, pieceI.width)
	assert.Equal(t, 3, pieceI.height)
	assertBoardStateWithAI(t, board.fields, pieceI.x, pieceI.y)

	pieceI.Rotate(&board)
	assert.Equal(t, 270, pieceI.rotation)
	assert.Equal(t, 3, pieceI.width)
	assert.Equal(t, 1, pieceI.height)
	assert.True(t, board.fields[x][y])
	assert.True(t, board.fields[x+1][y])
	assert.True(t, board.fields[x+2][y])

	pieceI.Rotate(&board)
	assert.Equal(t, 0, pieceI.rotation)
	assert.Equal(t, 1, pieceI.width)
	assert.Equal(t, 3, pieceI.height)
	assertBoardStateWithAI(t, board.fields, pieceI.x, pieceI.y)
}

func TestPieceRotationWhenItIsNotFailing(t *testing.T) {
	board := NewBoard(pieceT)
	piece := boardPiece{isFalling: false}

	rotation, err := piece.Rotate(&board)

	assert.Equal(t, errors.New("cannot rotate a stopped piece"), err)
	assert.Equal(t, 0, rotation)
}

func TestThatPieceDoesNotRotateWhenTooCloseFromRightBoard(t *testing.T) {
	board := NewBoard(pieceI)
	piece := &board.fallingPiece

	for i := 0; i < 22; i++ {
		piece.MoveRight(&board)
	}

	rotation, err := piece.Rotate(&board)

	assert.Equal(t, 0, rotation)
	assert.Equal(t, errors.New("is impossible to rotate in this position"), err)
}

func TestThatIPieceDoesNotRotateWhenTooCloseFromOtherPieces(t *testing.T) {
	board := NewBoard(pieceI)
	piece := &board.fallingPiece

	for j := 0; j < 28; j++ {
		board.fields[26][j] = true
	}

	rotation, err := piece.Rotate(&board)

	assert.Equal(t, 0, rotation)
	assert.Equal(t, errors.New("is impossible to rotate in this position"), err)
}

func TestThatLPieceDoesNotRotateWhenTooCloseFromOtherPieces(t *testing.T) {
	board := NewBoard(pieceL)
	piece := &board.fallingPiece

	for j := 0; j < 28; j++ {
		board.fields[26][j] = true
	}

	rotation, err := piece.Rotate(&board)

	assert.Equal(t, 0, rotation)
	assert.Equal(t, errors.New("is impossible to rotate in this position"), err)
}

func TestThatTPieceDoesNotRotateWhenTooCloseFromOtherPieces(t *testing.T) {
	board := NewBoard(pieceT)
	piece := &board.fallingPiece

	for j := 0; j < 28; j++ {
		board.fields[26][j] = true
	}
	piece.Rotate(&board)
	piece.MoveRight(&board)

	rotation, err := piece.Rotate(&board)

	assert.Equal(t, 90, rotation)
	assert.Equal(t, errors.New("is impossible to rotate in this position"), err)
}

func TestThatSPieceDoesNotRotateWhenTooCloseFromOtherPieces(t *testing.T) {
	board := NewBoard(pieceS)
	piece := &board.fallingPiece

	for j := 0; j < 28; j++ {
		board.fields[26][j] = true
	}

	rotation, err := piece.Rotate(&board)

	assert.Equal(t, 90, rotation)
	assert.Nil(t, err)
}

func TestThatDotPieceDoesNotRotateWhenTooCloseFromOtherPieces(t *testing.T) {
	board := NewBoard(pieceDot)
	piece := &board.fallingPiece

	for j := 0; j < 28; j++ {
		board.fields[26][j] = true
	}

	piece.MoveRight(&board)

	rotation, err := piece.Rotate(&board)

	assert.Equal(t, 90, rotation)
	assert.Nil(t, err)
}

func TestThatFallingPieceMovesLeft(t *testing.T) {
	board := NewBoard(pieceL)
	steps := 10
	piece := board.GetFallingPiece()

	assert.Equal(t, 24, piece.x)

	piece.MoveLeft(&board)
	piece.MoveLeft(&board)
	piece.MoveLeft(&board)

	for i := 0; i < steps; i++ {
		board.Tick()
	}

	assert.Equal(t, 21, piece.x)
	assert.Equal(t, 10, piece.y)
}

func TestThatFallingPieceMovesLeftUntilTheBorder(t *testing.T) {
	board := NewBoard(pieceL)
	steps := 50
	piece := board.GetFallingPiece()

	for i := 0; i < steps; i++ {
		piece.MoveLeft(&board)
	}

	assert.Equal(t, 0, piece.x)
	assert.Equal(t, 0, piece.y)
}

func TestThatFallingPieceMovesRight(t *testing.T) {
	board := NewBoard(pieceL)
	steps := 10
	piece := board.GetFallingPiece()

	piece.MoveRight(&board)
	piece.MoveRight(&board)
	piece.MoveRight(&board)

	for i := 0; i < steps; i++ {
		board.Tick()
	}

	assert.Equal(t, 27, piece.x)
	assert.Equal(t, 10, piece.y)
}

func TestThatFallingPieceMovesRightUntilTheBorder(t *testing.T) {
	board := NewBoard(pieceL)
	steps := 50
	piece := board.GetFallingPiece()

	for i := 0; i < steps; i++ {
		piece.MoveRight(&board)
	}

	assert.Equal(t, len(board.fields)-piece.width, piece.x)
	assert.Equal(t, 0, piece.y)
	assertBoardStateWithAnL(t, board.fields, piece.x, piece.y)
}

func TestThatDotPieceDoesNotMoveRightWhenTooCloseFromOtherPieces(t *testing.T) {
	board := NewBoard(pieceDot)
	piece := &board.fallingPiece
	expectedX := piece.x + 1
	for j := 0; j < 28; j++ {
		board.fields[26][j] = true
	}

	piece.MoveRight(&board)

	x, err := piece.MoveRight(&board)

	assert.Equal(t, expectedX, x)
	assert.Equal(t, errors.New("is impossible to move right from this position"), err)
}

func TestThatLPieceDoesNotMoveRightWhenTooCloseFromOtherPieces(t *testing.T) {
	board := NewBoard(pieceL)
	piece := &board.fallingPiece
	expectedX := piece.x - 1
	for j := 0; j < 28; j++ {
		board.fields[22][j] = true
	}

	piece.MoveLeft(&board)

	x, err := piece.MoveLeft(&board)

	assert.Equal(t, expectedX, x)
	assert.Equal(t, errors.New("is impossible to move left from this position"), err)
}
