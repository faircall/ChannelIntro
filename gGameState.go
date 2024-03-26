package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameStateT int

const (
	OpeningT GameStateT = iota // this is 0
	CasinoT
	NumGameStatesT
)

type GameState struct {
	CurrentState GameStateT
	CurrentTime  float64
	DeltaTime    float32
	TimeElapsed  float32
	Scenes       []Scene
	Camera       Camera
}

type GameUpdateFunction func(gameState *GameState)

func Update(gameState *GameState, gameUpdateFunctions []GameUpdateFunction) {
	gameUpdateFunctions[gameState.CurrentState](gameState)
}

func OpeningUpdate(gameState *GameState) {
	gameState.CurrentTime = rl.GetTime()
	gameState.DeltaTime = rl.GetFrameTime()
	gameState.TimeElapsed += gameState.DeltaTime
}

func CasinoUpdate(gameState *GameState) {
	gameState.CurrentTime = rl.GetTime()
	gameState.DeltaTime = rl.GetFrameTime()
	gameState.TimeElapsed += gameState.DeltaTime

	if rl.IsKeyPressed(rl.KeyF1) {
		if rl.IsWindowFullscreen() {
			rl.SetWindowState(rl.FlagWindowResizable)
		} else {
			rl.SetWindowState(rl.FlagFullscreenMode)
		}
	}
}
