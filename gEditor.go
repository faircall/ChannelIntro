package main

import rl "github.com/gen2brain/raylib-go/raylib"

func Vec2InRect(v rl.Vector2, r rl.Rectangle) bool {
	if v.X >= r.X && v.X <= (r.X+r.Width) &&
		v.Y >= r.Y && v.Y <= (r.Y+r.Height) {
		return true
	}
	return false
}

func UpdateEditor(gameState *GameState) {
	// get input, which will be mouse position, keys pressed
	mousePos := rl.GetMousePosition()
	gameState.MousePos = mousePos

	switch gameState.Editor.CurrentState {
	case InitialT:
		// new or load existing
		for i, rect := range gameState.Editor.Rects {
			if Vec2InRect(mousePos, rect.Rect) {
				rect.IsSelected = true
				rect.ColorLerp = min(rect.ColorLerp+gameState.DeltaTime*2, 1.0)
				gameState.Editor.Rects[i] = rect
			} else {
				rect.IsSelected = false
				rect.ColorLerp = max(rect.ColorLerp-gameState.DeltaTime*2, 0.0)
				gameState.Editor.Rects[i] = rect
			}
		}
		// if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {

		// }
	}
}

func DrawRect(rect rl.Rectangle, color rl.Color) {
	rl.DrawRectangle(int32(rect.X), int32(rect.Y), int32(rect.Width), int32(rect.Height), color)
}

func LerpColors(colStart rl.Color, colEnd rl.Color, lerp float32) rl.Color {
	rDiff := int32(colEnd.R) - int32(colStart.R)
	gDiff := int32(colEnd.G) - int32(colStart.G)
	bDiff := int32(colEnd.B) - int32(colStart.B)

	result := rl.Color{R: colStart.R + uint8(float32(rDiff)*lerp), G: colStart.G + uint8(float32(gDiff)*lerp), B: colStart.B + uint8(float32(bDiff)*lerp), A: 255}
	return result
}

func RenderEditor(gameState GameState) {
	rl.ClearBackground(rl.Red)

	rl.DrawCircle(int32(gameState.MousePos.X), int32(gameState.MousePos.Y), 10.0, rl.White)

	switch gameState.Editor.CurrentState {
	case InitialT:
		for _, rect := range gameState.Editor.Rects {
			if rect.IsSelected {
				drawColor := LerpColors(rect.BaseColor, rect.HighlightColor, rect.ColorLerp)
				DrawRect(rect.Rect, drawColor)
				rl.DrawText(rect.Text, int32(rect.Rect.X+rect.Rect.Width/2), int32(rect.Rect.Y+rect.Rect.Height/2), 18, rect.HighlightTextColor)
			} else {
				drawColor := LerpColors(rect.BaseColor, rect.HighlightColor, rect.ColorLerp)
				DrawRect(rect.Rect, drawColor)
				rl.DrawText(rect.Text, int32(rect.Rect.X+rect.Rect.Width/2), int32(rect.Rect.Y+rect.Rect.Height/2), 16, rect.BaseTextColor)
			}
		}

	}
}

func SaveState() {

}
