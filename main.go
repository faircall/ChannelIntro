package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1280, 720, "Bond Channel Intro")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	var gameState GameState
	var testScene Scene

	var dotGroup1 DotGroup
	dotGroup1.Dots = MakeRectangularGroupOfCircles(30, 30, 10.0, 10, 10, rl.Red, 20.0)
	dotGroup1.StartTime = 0.0
	dotGroup1.EndTime = 4.0
	testScene.Dots = append(testScene.Dots, dotGroup1)

	var dotGroup2 DotGroup
	dotGroup2.Dots = MakeRectangularGroupOfCircles(540, 70, 10.0, 4, 4, rl.Green, 40.0)
	dotGroup2.StartTime = 5.0
	dotGroup2.EndTime = 10.0
	testScene.Dots = append(testScene.Dots, dotGroup2)

	gameState.Scenes = append(gameState.Scenes, testScene)
	gameState.Camera = Camera{Position: rl.Vector2{X: 0.0, Y: 0.0}, Color: rl.RayWhite, Focus: 1.0}
	var renderFunctions = []RenderFunction{RenderOpening}
	var updateFunctions = []GameUpdateFunction{OpeningUpdate}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		Update(&gameState, updateFunctions)
		Render(gameState, renderFunctions)

		rl.EndDrawing()
	}
}
