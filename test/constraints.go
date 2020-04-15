package main

import (
	rl "github.com/Lachee/raylib-goplus/raylib"
	"zgui"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)

	rl.InitWindow(800, 450, "raygui test")

	display := zgui.GetDisplay()

	constraints := zgui.DefaultConstraints()
	constraints.SetX(zgui.NewPixelConstraint(100))
	constraints.SetY(zgui.NewPixelConstraint(200))
	constraints.SetWidth(zgui.NewPixelConstraint(100))
	constraints.SetHeight(zgui.NewPixelConstraint(200))
	box := zgui.NewBoxComponent(rl.Red)

	display.Add(box, constraints)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		display.Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
