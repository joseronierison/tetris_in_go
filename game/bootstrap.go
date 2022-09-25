package game

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func Start(ss *ScreenState) {
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
		case MainMenuScreen:
			InitMenu(ss)
		case GameScreen:
			InitGameBoard(ss)
		case Exit:
			break mainloop
		}
	}
}
