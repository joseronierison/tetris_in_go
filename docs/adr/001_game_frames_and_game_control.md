# Board Data Structure

## Status

This version was accepted for the first version 0.0.0

## Context

As a game need to change it status time by time as the actions and rules are applied and the users controls the game.


## Decision

A time tick was used to iterate over a defined period of time where the games will change when each time the `board.Tick()` method is called.

```golang
func drawBoardFrames() {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

    // [..]

	for {
		drawGameBoard()

		drawScore()
		drawNextPiece()

		tetrisBoard.Tick()

		// ..
		<-ticker.C
	}
}
```

This method is called in a go thread different of the one who listen to the user iteraction.

## Consequences

It enables the changes of board state in a easy way. A possible challenge is to acelerate the game frames as the rules of the game requires that.
For instance, when the user hits the arrow down button we could acelerate the falling piece movement as a useful controle of the game. It could represente a chalenge since we need to keep that behavior only while the key is pressed and go back to the normal aceleration.
