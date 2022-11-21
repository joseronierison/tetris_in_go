package core

type board struct {
	fields       [49][28]bool
	fallingPiece boardPiece
	nextPiece    boardPiece
}

func NewBoard(startPiece boardPiece) board {
	var fields [49][28]bool

	for x := 0; x < 49; x++ {
		for y := 0; y < 28; y++ {
			fields[x][y] = false
		}
	}

	board := board{
		fields:       fields,
		fallingPiece: startPiece,
		nextPiece:    GenerateRandomFallingPiece(),
	}

	board.place(startPiece)
	return board
}

func (board *board) Tick() {
	if board.isNextFieldOccupied(board.fallingPiece) {
		board.fallingPiece.isFalling = false
		board.fallingPiece = board.nextPiece
		board.place(board.fallingPiece)
		board.nextPiece = GenerateRandomFallingPiece()
	}

	if board.fallingPiece.y >= len(board.fields[0])-board.fallingPiece.height {
		board.fallingPiece.isFalling = false
		board.fallingPiece = board.nextPiece
		board.place(board.fallingPiece)
		board.nextPiece = GenerateRandomFallingPiece()
		return
	}

	board.remove(board.fallingPiece)
	board.fallingPiece.y++
	board.place(board.fallingPiece)
}

func (board *board) isNextFieldOccupied(boardPiece boardPiece) bool {
	nextColumn := boardPiece.y + boardPiece.height
	if nextColumn < len(board.fields[0]) {
		return board.fields[boardPiece.x][nextColumn]
	}

	return false
}

func (board *board) remove(boardPiece boardPiece) {
	board.setStateOnBoard(boardPiece, false)
}

func (board *board) place(boardPiece boardPiece) {
	board.setStateOnBoard(boardPiece, true)
}

func (board *board) setStateOnBoard(boardPiece boardPiece, state bool) {
	x := boardPiece.x
	y := boardPiece.y

	switch boardPiece.pieceType {
	case 'T':
		if boardPiece.rotation == 0 {
			board.fields[x][y] = state
			board.fields[x+1][y] = state
			board.fields[x+1][y+1] = state
			board.fields[x+2][y] = state
		} else if boardPiece.rotation == 90 {
			board.fields[x][y] = state
			board.fields[x][y+1] = state
			board.fields[x+1][y+1] = state
			board.fields[x][y+2] = state
		} else if boardPiece.rotation == 180 {
			board.fields[x][y+2] = state
			board.fields[x+1][y+2] = state
			board.fields[x+1][y+1] = state
			board.fields[x+2][y+2] = state
		} else if boardPiece.rotation == 270 {
			board.fields[x+2][y] = state
			board.fields[x+2][y+1] = state
			board.fields[x+1][y+1] = state
			board.fields[x+2][y+2] = state
		}
	case 'L':
		if boardPiece.rotation == 0 {
			board.fields[x][y] = state
			board.fields[x][y+1] = state
			board.fields[x][y+2] = state
			board.fields[x+1][y+2] = state
			board.fields[x+2][y+2] = state
		} else if boardPiece.rotation == 90 {
			board.fields[x][y+2] = state
			board.fields[x+1][y+2] = state
			board.fields[x+2][y+2] = state
			board.fields[x+2][y+1] = state
			board.fields[x+2][y] = state
		} else if boardPiece.rotation == 180 {
			board.fields[x][y] = state
			board.fields[x+1][y] = state
			board.fields[x+2][y] = state
			board.fields[x+2][y+1] = state
			board.fields[x+2][y+2] = state
		} else if boardPiece.rotation == 270 {
			board.fields[x][y] = state
			board.fields[x+1][y] = state
			board.fields[x+2][y] = state
			board.fields[x][y+1] = state
			board.fields[x][y+2] = state
		}
	case 'I':
		if boardPiece.rotation == 0 || boardPiece.rotation == 180 {
			board.fields[x][y] = state
			board.fields[x][y+1] = state
			board.fields[x][y+2] = state
		} else if boardPiece.rotation == 90 || boardPiece.rotation == 270 {
			board.fields[x][y+2] = state
			board.fields[x+1][y+2] = state
			board.fields[x+2][y+2] = state
		}
	case 'S':
		board.fields[x][y] = state
		board.fields[x][y+1] = state
		board.fields[x+1][y] = state
		board.fields[x+1][y+1] = state
	case '.':
		board.fields[x][y] = state
	}
}

func (board *board) GetFallingPiece() *boardPiece {
	return &board.fallingPiece
}

func (board *board) GetNexPiece() *boardPiece {
	return &board.nextPiece
}

func (board *board) GetFields() [49][28]bool {
	return board.fields
}
