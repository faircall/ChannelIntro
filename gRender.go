package main

import rl "github.com/gen2brain/raylib-go/raylib"

type RenderFunction func(gameState GameState)

func Render(gameState GameState, renderFunctions []RenderFunction) {
	renderFunctions[gameState.CurrentState](gameState)
}

func RenderOpening(gameState GameState) {
	rl.ClearBackground(rl.Pink)
}
