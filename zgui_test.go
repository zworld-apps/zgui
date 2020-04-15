package zgui

import (
	"fmt"
	"testing"

	rl "github.com/Lachee/raylib-goplus/raylib"
)

func TestBoxComponent(t *testing.T) {
	defer rl.CloseWindow()

	display := GetDisplay()

	constraints := DefaultConstraints()
	box := NewBoxComponent(rl.Red)

	display.Add(box, constraints)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		//rl.ClearBackground(rl.RayWhite)

		display.Draw()

		rl.EndDrawing()
	}

	fmt.Println("should close")
}

func TestMain(m *testing.M) {
	rl.InitWindow(800, 450, "raygui test")
	rl.SetTargetFPS(60)

	m.Run()
}
