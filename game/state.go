package game

const MainMenuScreen = "main_menu"
const GameScreen = "game"
const Exit = "exit"

var SS = ScreenState{currentScreen: MainMenuScreen}

type ScreenState struct {
	currentScreen string
}

func (ss *ScreenState) GoToMainMenu() {
	ss.currentScreen = MainMenuScreen
}

func (ss *ScreenState) GoToGame() {
	ss.currentScreen = GameScreen
}

func (ss *ScreenState) Exit() {
	ss.currentScreen = Exit
}

func (ss *ScreenState) GetCurrentScreen() string {
	return ss.currentScreen
}
