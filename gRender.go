package main

import rl "github.com/gen2brain/raylib-go/raylib"

type RenderFunction func(gameState GameState)

type DrawableCircle struct {
	Position rl.Vector2
	Color    rl.Color
	Radius   float32
}

type Camera struct {
	Position rl.Vector2
	Color    rl.Color
	Focus    float64
	Width    float32
	Height   float32
}

type DotGroup struct {
	Dots      []DrawableCircle
	StartTime float32
	EndTime   float32
}

type TileMap struct {
	Tiles      []int
	Width      int
	Height     int
	TileWidth  int
	TileHeight int
}

type Scene struct {
	Dots    []DotGroup
	TileMap []TileMap
}

func Render(gameState GameState, renderFunctions []RenderFunction) {
	renderFunctions[gameState.CurrentState](gameState)
}

func DrawCircles(drawables []DrawableCircle, camera Camera) {
	for _, circ := range drawables {
		// do we actually need the radius to be dependent on screenwidth and height?
		rl.DrawCircle(int32((circ.Position.X-camera.Position.X)*camera.Width), int32((circ.Position.Y-camera.Position.Y)*camera.Height), circ.Radius*camera.Height, circ.Color)
	}
}

func RenderOpening(gameState GameState) {
	rl.ClearBackground(rl.Black)
	for _, dots := range gameState.Scenes[gameState.CurrentState].Dots {
		if gameState.TimeElapsed >= dots.StartTime && gameState.TimeElapsed <= dots.EndTime {
			DrawCircles(dots.Dots, gameState.Camera)
		}
	}
}

func DrawRectangle(x int, y int, width int, height int, camera Camera, color rl.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), color)
}

func RenderCasino(gameState GameState) {
	rl.ClearBackground(rl.Pink)
	tileWidth := gameState.Scenes[gameState.CurrentState].TileMap[0].TileWidth
	for y := 0; y < gameState.Scenes[gameState.CurrentState].TileMap[0].Height; y++ {
		for x := 0; x < gameState.Scenes[gameState.CurrentState].TileMap[0].Width; x++ {
			if (y*x+1)%2 == 0 {
				DrawRectangle(x*tileWidth, y*tileWidth, tileWidth, tileWidth, gameState.Camera, rl.Blue)
			}
			if (y*x+1)%3 == 0 {
				DrawRectangle(x*tileWidth, y*tileWidth, tileWidth, tileWidth, gameState.Camera, rl.Green)
			}
			if (y*x+1)%5 == 0 {
				DrawRectangle(x*tileWidth, y*tileWidth, tileWidth, tileWidth, gameState.Camera, rl.Red)
			}
			if (y*x+1)%7 == 0 {
				DrawRectangle(x*tileWidth, y*tileWidth, tileWidth, tileWidth, gameState.Camera, rl.Yellow)
			}
		}
	}

}

func MakeRectangularGroupOfCircles(startX float32, startY float32, spaceX float32, spaceY float32, across int, down int, color rl.Color, radius float32) []DrawableCircle {
	var result []DrawableCircle
	for y := 0; y < down; y++ {
		for x := 0; x < across; x++ {
			position := rl.Vector2{X: startX + float32(x)*(radius*2.0+spaceX), Y: startY + float32(y)*(radius*2.0+spaceY)}
			circToAdd := DrawableCircle{Position: position, Color: color, Radius: radius}
			result = append(result, circToAdd)
		}
	}
	return result
}
