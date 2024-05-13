package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	screenWidth := 960
	screenHeight := 540
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Bond Channel Intro")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	var gameState GameState

	var testScene = MakeOpeningData()
	var casinoScene = MakeCasinoData()
	gameState.Scenes = append(gameState.Scenes, testScene)
	gameState.Scenes = append(gameState.Scenes, casinoScene)
	gameState.ScreenWidth = int32(screenWidth)
	gameState.ScreenHeight = int32(screenHeight)
	gameState.Camera = Camera{Position: rl.Vector2{X: 0.0, Y: 0.0}, Color: rl.RayWhite, Focus: 1.0, Width: float32(screenWidth), Height: float32(screenHeight)}
	var renderFunctions = []RenderFunction{RenderOpening, RenderCasino, RenderEditor}
	var updateFunctions = []GameUpdateFunction{OpeningUpdate, CasinoUpdate, UpdateEditor}

	gameState.CurrentState = EditorT
	gameState.Editor.Rects = append(gameState.Editor.Rects, SelectableRect{Rect: rl.Rectangle{X: 0, Y: 0, Width: float32(screenWidth / 2), Height: float32(screenHeight)}, IsSelected: false, BaseColor: rl.Blue, BaseTextColor: rl.Red, HighlightColor: rl.Green, HighlightTextColor: rl.Black, Text: "NEW MAP"})
	gameState.Editor.Rects = append(gameState.Editor.Rects, SelectableRect{Rect: rl.Rectangle{X: float32(screenWidth / 2), Y: 0, Width: float32(screenWidth / 2), Height: float32(screenHeight)}, IsSelected: false, BaseColor: rl.Red, BaseTextColor: rl.Blue, HighlightColor: rl.Green, HighlightTextColor: rl.Black, Text: "LOAD MAP"})

	for !rl.WindowShouldClose() {

		Update(&gameState, updateFunctions)
		rl.BeginDrawing()
		Render(gameState, renderFunctions)
		rl.EndDrawing()

	}
}
