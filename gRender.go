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
}

type DotGroup struct {
	Dots      []DrawableCircle
	StartTime float32
	EndTime   float32
}

type Scene struct {
	Dots []DotGroup
}

func Render(gameState GameState, renderFunctions []RenderFunction) {
	renderFunctions[gameState.CurrentState](gameState)
}

func DrawCircles(drawables []DrawableCircle, camera Camera) {
	for _, circ := range drawables {
		rl.DrawCircle(int32(circ.Position.X-camera.Position.X), int32(circ.Position.Y-camera.Position.Y), circ.Radius, circ.Color)
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

func MakeRectangularGroupOfCircles(startX int, startY int, space float32, across int, down int, color rl.Color, radius float32) []DrawableCircle {
	var result []DrawableCircle
	for y := 0; y < down; y++ {
		for x := 0; x < across; x++ {
			position := rl.Vector2{X: float32(startX) + float32(x)*(radius*2.0+space), Y: float32(startY) + float32(y)*(radius*2.0+space)}
			circToAdd := DrawableCircle{Position: position, Color: color, Radius: radius}
			result = append(result, circToAdd)
		}
	}
	return result
}
