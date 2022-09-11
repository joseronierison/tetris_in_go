package game

import (
	"fmt"
	"os"
	"tetris/screens"
	gameboard "tetris/screens/gameBoard"
	"tetris/screens/menu"

	"github.com/nsf/termbox-go"
)

func Start(ss *screens.ScreenState) {
	err := termbox.Init()
	termbox.SetInputMode(termbox.InputEsc)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer termbox.Close()

	for {
		switch ss.GetCurrentScreen() {
		case screens.MainMenuScreen:
			menu.Init()
		case screens.GameScreen:
			gameboard.Init()
		case "exit":
			return
		}
	}
}
