package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameStateT int
type EditorStateT int

const (
	OpeningT GameStateT = iota // this is 0
	CasinoT
	EditorT
	NumGameStatesT
)

const (
	InitialT EditorStateT = iota // this is 0
	NumEditorStatesT
)

type GameState struct {
	CurrentState GameStateT
	CurrentTime  float64
	DeltaTime    float32
	TimeElapsed  float32
	Scenes       []Scene
	Camera       Camera
	MousePos     rl.Vector2
	Editor       EditorState
	ScreenWidth  int32
	ScreenHeight int32
}

type EditorState struct {
	CurrentState     EditorStateT
	CurrentSelection int32
	Rects            []SelectableRect
}

type SelectableRect struct {
	Rect               rl.Rectangle
	IsSelected         bool
	BaseColor          rl.Color
	BaseTextColor      rl.Color
	HighlightColor     rl.Color
	HighlightTextColor rl.Color
	Text               string
	ColorLerp          float32
}

type GameUpdateFunction func(gameState *GameState)

func Update(gameState *GameState, gameUpdateFunctions []GameUpdateFunction) {
	gameState.DeltaTime = rl.GetFrameTime()
	gameState.CurrentTime = rl.GetTime()
	gameState.TimeElapsed += gameState.DeltaTime
	gameUpdateFunctions[gameState.CurrentState](gameState)
}

func OpeningUpdate(gameState *GameState) {

}

func CasinoUpdate(gameState *GameState) {
	if rl.IsKeyPressed(rl.KeyF1) {
		if rl.IsWindowFullscreen() {
			rl.SetWindowState(rl.FlagWindowResizable)
		} else {
			rl.SetWindowState(rl.FlagFullscreenMode)
		}
	}
}
