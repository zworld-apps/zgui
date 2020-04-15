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
	constraints.SetX(zgui.NewCenterConstraint())
	constraints.SetY(zgui.NewPixelConstraint(20))
	constraints.SetWidth(zgui.NewRelativeConstraint(0.1))
	constraints.SetHeight(zgui.NewAspectConstraint(1))
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
