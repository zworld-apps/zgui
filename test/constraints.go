package main

import (
	"fmt"
	"zgui"

	rl "github.com/xzebra/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)

	rl.InitWindow(800, 450, "raygui test")

	display := zgui.GetDisplay()

	constraints := zgui.DefaultConstraints()
	constraints.SetX(zgui.NewRelativeConstraint(0.3))
	constraints.SetY(zgui.NewFillConstraint())
	constraints.SetWidth(zgui.NewRelativeConstraint(0.7))
	constraints.SetHeight(zgui.NewFillConstraint())
	redBox := zgui.NewBoxComponent(&zgui.BoxOptions{
		Color:     rl.Red,
		Roundness: 0.0,
		Segments:  5,
	})

	display.Add(redBox, constraints)

	constraints = zgui.DefaultConstraints()
	constraints.SetX(zgui.NewCenterConstraint())
	constraints.SetWidth(zgui.NewRelativeConstraint(0.3))
	constraints.SetHeight(zgui.NewPixelConstraint(20))
	textField := zgui.NewTextFieldComponent(&zgui.TextFieldOptions{
		Box: &zgui.BoxOptions{
			Color:     rl.Black,
			Roundness: 0.3,
			Segments:  8,
		},
		Text: &zgui.TextOptions{
			Color: rl.Black,
		},
		SubmitCallback: func(tf *zgui.TextFieldComponent) {
			fmt.Println(tf.Label.Text)
		},
	})

	redBox.Add(textField, constraints)

	constraints = zgui.DefaultConstraints()
	constraints.SetY(zgui.NewCenterConstraint())
	constraints.SetWidth(zgui.NewRelativeConstraint(0.3))
	greenBox := zgui.NewBoxComponent(&zgui.BoxOptions{
		Color:     rl.Green,
		Roundness: 0.1,
		Segments:  20,
	})

	display.Add(greenBox, constraints)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		display.Update(rl.GetFrameTime())
		display.Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
