package core

import (
	"errors"
	"math/rand"
)

type boardPiece struct {
	pieceType rune
	x         int
	y         int
	rotation  int
	width     int
	height    int
	isFalling bool
}

func (piece *boardPiece) moveLeft(board *board) {
	if piece.x < 1 {
		return
	}

	board.remove(*piece)
	piece.x--
	board.place(*piece)
}

func (piece *boardPiece) moveRight(board *board) {
	if piece.x > 48-piece.width {
		return
	}

	board.remove(*piece)
	piece.x++
	board.place(*piece)
}

func (piece *boardPiece) rotate(board *board) (int, error) {
	if !piece.isFalling {
		return piece.rotation, errors.New("cannot rotate a stopped piece")
	}

	board.remove(*piece)

	width := piece.width
	height := piece.height
	piece.width = height
	piece.height = width

	if piece.rotation == 270 {
		piece.rotation = 0
		board.place(*piece)
		return piece.rotation, nil
	}

	piece.rotation += 90
	board.place(*piece)

	return piece.rotation, nil
}

func GenerateRandomFallingPiece() boardPiece {
	pieceT := boardPiece{
		pieceType: 'T',
		x:         24,
		y:         0,
		rotation:  0,
		width:     3,
		height:    2,
		isFalling: true,
	}

	pieceL := boardPiece{
		pieceType: 'L',
		x:         24,
		y:         0,
		rotation:  0,
		width:     3,
		height:    3,
		isFalling: true,
	}

	pieceI := boardPiece{
		pieceType: 'I',
		x:         24,
		y:         0,
		rotation:  0,
		width:     1,
		height:    3,
		isFalling: true,
	}

	pieceS := boardPiece{
		pieceType: 'S',
		x:         24,
		y:         0,
		rotation:  0,
		width:     2,
		height:    2,
		isFalling: true,
	}

	pieceDot := boardPiece{
		pieceType: '.',
		x:         24,
		y:         0,
		rotation:  0,
		width:     1,
		height:    1,
		isFalling: true,
	}

	pieces := []boardPiece{pieceT, pieceL, pieceI, pieceS, pieceDot}
	index := rand.Intn(4)
	return pieces[index]
}
