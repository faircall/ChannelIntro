package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1280, 720, "Bond Channel Intro")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	var gameState GameState
	var renderFunctions = []RenderFunction{RenderOpening}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		Update(gameState)
		Render(gameState, renderFunctions)

		rl.EndDrawing()
	}
}
