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
	if board.fallingPiece.y >= len(board.fields[0])-board.fallingPiece.height {
		board.fallingPiece.isFalling = false
		board.fallingPiece = board.nextPiece
		board.place(board.fallingPiece)
		board.nextPiece = GenerateRandomFallingPiece()
		return
	}

	prospectedPiece := board.fallingPiece
	prospectedPiece.y++

	if board.IsFieldOccupied(board.fallingPiece, prospectedPiece) {
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

func (board *board) IsFieldOccupied(oldPiece, newPiece boardPiece) bool {
	intermediateBoard := *board
	intermediateBoard.remove(oldPiece)

	newPos := newPiece.GetPieceAtomsPositions()
	for i := 0; i < len(newPos); i++ {
		position := newPos[i]

		if position.y >= 28 {
			return false
		}

		if intermediateBoard.fields[position.x][position.y] {
			return true
		}
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
	positions := boardPiece.GetPieceAtomsPositions()

	for i := 0; i < len(positions); i++ {
		position := positions[i]
		board.fields[position.x][position.y] = state
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
