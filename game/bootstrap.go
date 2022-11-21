package game

import (
	"fmt"
	"os"
	"tetris/game/screens"

	"github.com/nsf/termbox-go"
)

func Start() {
	ss := &screens.SS
	err := termbox.Init()
	termbox.SetInputMode(termbox.InputEsc)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer termbox.Close()

mainloop:
	for {
		switch ss.GetCurrentScreen() {
		case screens.MainMenuScreen:
			screens.InitMenu(ss)
		case screens.GameScreen:
			screens.InitGameBoard(ss)
		case screens.Exit:
			break mainloop
		}
	}
}
