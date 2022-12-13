package core

import (
	"errors"
	"math/rand"
)

type position struct {
	x int
	y int
}

type boardPiece struct {
	pieceType rune
	x         int
	y         int
	rotation  int
	width     int
	height    int
	isFalling bool
}

func (piece *boardPiece) MoveLeft(board *board) (int, error) {
	prospectedPiece := *piece
	prospectedPiece.x--

	if piece.x < 1 || board.IsFieldOccupied(prospectedPiece) {
		return piece.x, errors.New("is impossible to move left from this position")
	}

	board.remove(*piece)
	piece.x--
	board.place(*piece)
	return piece.x, nil
}

func (piece *boardPiece) MoveRight(board *board) (int, error) {
	prospectedPiece := *piece
	prospectedPiece.x++

	if piece.x > 48-piece.width || board.IsFieldOccupied(prospectedPiece) {
		return piece.x, errors.New("is impossible to move right from this position")
	}

	board.remove(*piece)
	piece.x++
	board.place(*piece)

	return piece.x, nil
}

func (piece *boardPiece) Rotate(board *board) (int, error) {
	if !piece.isFalling {
		return piece.rotation, errors.New("cannot rotate a stopped piece")
	}

	if piece.height+piece.x > 48 {
		return piece.rotation, errors.New("is impossible to rotate in this position")
	}

	prospectedPiece := *piece
	prospectedPiece.rotation = piece.nextRotation()

	if board.IsFieldOccupied(prospectedPiece) {
		return piece.rotation, errors.New("is impossible to rotate in this position")
	}

	board.remove(*piece)

	width := piece.width
	height := piece.height
	piece.width = height
	piece.height = width

	piece.rotation = piece.nextRotation()
	board.place(*piece)

	return piece.rotation, nil
}

func (piece *boardPiece) nextRotation() int {
	if piece.rotation == 270 {
		return 0
	}

	return piece.rotation + 90
}

func (piece *boardPiece) GetPieceType() rune {
	return piece.pieceType
}

func (piece *boardPiece) GetPieceAtomsPositions() []position {
	var pieceAtomsPositions []position

	switch piece.pieceType {
	case 'T':
		pieceAtomsPositions = piece.getTPieceAtomsPositions()
	case 'L':
		pieceAtomsPositions = piece.getLPieceAtomsPositions()
	case 'I':
		pieceAtomsPositions = piece.getIPieceAtomsPositions()
	case 'S':
		pieceAtomsPositions = piece.getSPieceAtomsPositions()
	case '.':
		pieceAtomsPositions = piece.getDotPieceAtomsPositions()
	}

	return pieceAtomsPositions
}

func (piece *boardPiece) getTPieceAtomsPositions() []position {
	if piece.rotation == 90 {
		return []position{
			{x: piece.x, y: piece.y},
			{x: piece.x, y: piece.y + 1},
			{x: piece.x + 1, y: piece.y + 1},
			{x: piece.x, y: piece.y + 2},
		}
	} else if piece.rotation == 180 {
		return []position{
			{x: piece.x, y: piece.y + 1},
			{x: piece.x + 1, y: piece.y + 1},
			{x: piece.x + 1, y: piece.y},
			{x: piece.x + 2, y: piece.y + 1},
		}
	} else if piece.rotation == 270 {
		return []position{
			{x: piece.x + 1, y: piece.y},
			{x: piece.x + 1, y: piece.y + 1},
			{x: piece.x, y: piece.y + 1},
			{x: piece.x + 1, y: piece.y + 2},
		}
	}

	return []position{
		{x: piece.x, y: piece.y},
		{x: piece.x + 1, y: piece.y},
		{x: piece.x + 1, y: piece.y + 1},
		{x: piece.x + 2, y: piece.y},
	}
}

func (piece *boardPiece) getLPieceAtomsPositions() []position {
	if piece.rotation == 90 {
		return []position{
			{x: piece.x, y: piece.y + 2},
			{x: piece.x + 1, y: piece.y + 2},
			{x: piece.x + 2, y: piece.y + 2},
			{x: piece.x + 2, y: piece.y + 1},
			{x: piece.x + 2, y: piece.y},
		}
	} else if piece.rotation == 180 {
		return []position{
			{x: piece.x, y: piece.y},
			{x: piece.x + 1, y: piece.y},
			{x: piece.x + 2, y: piece.y},
			{x: piece.x + 2, y: piece.y + 1},
			{x: piece.x + 2, y: piece.y + 2},
		}
	} else if piece.rotation == 270 {
		return []position{
			{x: piece.x, y: piece.y},
			{x: piece.x + 1, y: piece.y},
			{x: piece.x + 2, y: piece.y},
			{x: piece.x, y: piece.y + 1},
			{x: piece.x, y: piece.y + 2},
		}
	}

	return []position{
		{x: piece.x, y: piece.y},
		{x: piece.x, y: piece.y + 1},
		{x: piece.x, y: piece.y + 2},
		{x: piece.x + 1, y: piece.y + 2},
		{x: piece.x + 2, y: piece.y + 2},
	}
}

func (piece *boardPiece) getIPieceAtomsPositions() []position {
	if piece.rotation == 90 || piece.rotation == 270 {
		return []position{
			{x: piece.x, y: piece.y},
			{x: piece.x + 1, y: piece.y},
			{x: piece.x + 2, y: piece.y},
		}
	}

	return []position{
		{x: piece.x, y: piece.y},
		{x: piece.x, y: piece.y + 1},
		{x: piece.x, y: piece.y + 2},
	}
}

func (piece *boardPiece) getSPieceAtomsPositions() []position {
	return []position{
		{x: piece.x, y: piece.y},
		{x: piece.x, y: piece.y + 1},
		{x: piece.x + 1, y: piece.y},
		{x: piece.x + 1, y: piece.y + 1},
	}
}

func (piece *boardPiece) getDotPieceAtomsPositions() []position {
	return []position{
		{x: piece.x, y: piece.y},
	}
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
