package core

type board struct {
	fields       [49][28]bool
	fallingPiece boardPiece
	nextPiece    boardPiece
	score        int
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
	if board.IsOver() {
		return
	}

	prospectedPiece := board.fallingPiece
	prospectedPiece.y++

	if board.hasAVerticalColisionForFallingPiece() {
		board.fallingPiece.isFalling = false
		board.computeScore()
		board.fallingPiece = board.nextPiece
		board.place(board.fallingPiece)
		board.nextPiece = GenerateRandomFallingPiece()
		return
	}

	board.remove(board.fallingPiece)
	board.fallingPiece.y++
	board.place(board.fallingPiece)
}

func (board *board) GetScore() int {
	return board.score
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

func (board *board) IsFieldOccupied(prospectedPiece boardPiece) bool {
	intermediateBoard := *board
	intermediateBoard.remove(board.fallingPiece)

	newPos := prospectedPiece.GetPieceAtomsPositions()
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

func (board *board) hasAVerticalColisionForFallingPiece() bool {
	prospectedPiece := board.fallingPiece
	prospectedPiece.y++

	return board.IsFieldOccupied(prospectedPiece) || board.isPieceOnBoardLimit()
}

func (board *board) isPieceOnBoardLimit() bool {
	return board.fallingPiece.y >= len(board.fields[0])-board.fallingPiece.height
}

func (board *board) IsOver() bool {
	return board.hasAVerticalColisionForFallingPiece() && board.fallingPiece.y == 0
}

func (board *board) computeScore() {
	var lines [28][49]bool = board.getInvertedBoardFields()
	var cl = 27

	for cl >= 0 {
		line := lines[cl]

		if isAll(line, true) {
			board.wipeOutBoardLine(cl)
			board.score += 10
			board.rearrange()
		}

		cl--
	}
}

func (board *board) rearrange() {
	for j := 27; j > 0; j-- {
		for i := 0; i < 49; i++ {
			board.fields[i][j] = board.fields[i][j-1]
		}
	}

}

func isAll(line [49]bool, status bool) bool {
	for i := 0; i < 49; i++ {
		if line[i] != status {
			return false
		}
	}

	return true
}

func (board *board) getInvertedBoardFields() [28][49]bool {
	var lines [28][49]bool
	for i := 0; i < 49; i++ {
		for j := 0; j < 28; j++ {
			lines[j][i] = board.fields[i][j]
		}
	}

	return lines
}

func (board *board) wipeOutBoardLine(lineNumber int) {
	for i := 0; i < 49; i++ {
		board.fields[i][lineNumber] = false
	}
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
