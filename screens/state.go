package screens

const MainMenuScreen = "main_menu"
const GameScreen = "game"

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

func (ss *ScreenState) GetCurrentScreen() string {
	return ss.currentScreen
}
