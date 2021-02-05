package zgui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// displayPadding is the separation from the OS window borders
const displayPadding = 10

// holdInsideWindow holds the mouse inside window frame.
func holdInsideWindow(mouse rl.Vector2) rl.Vector2 {
	sw, sh := float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())
	if mouse.X > (sw - displayPadding) {
		mouse.X = sw - displayPadding
	} else if mouse.X < 0 {
		mouse.X = displayPadding
	}

	if mouse.Y > (sh - displayPadding) {
		mouse.Y = sh - displayPadding
	} else if mouse.Y < 0 {
		mouse.Y = displayPadding
	}

	return mouse
}

func clamp(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
