package main

import (
	"fmt"
	"zgui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)

	rl.InitWindow(800, 450, "zgui test")

	zgui.LoadDefaultStyle()

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

	window := zgui.NewWindowComponent(&zgui.WindowOptions{
		Bar: &zgui.BoxOptions{
			Color:     rl.Blue,
			Roundness: 0.3,
			Segments:  8,
		},
		CloseBtn: &zgui.TextOptions{
			Color: rl.White,
		},
		Content: &zgui.BoxOptions{
			Color:     rl.White,
			Roundness: 0,
			Segments:  0,
		},
	})

	constraints = zgui.DefaultConstraints()
	constraints.SetX(zgui.NewPixelConstraint(200))
	constraints.SetY(zgui.NewPixelConstraint(50))
	constraints.SetWidth(zgui.NewPixelConstraint(200))
	constraints.SetHeight(zgui.NewPixelConstraint(300))
	display.Add(window, constraints)

	// PANEL
	panel := zgui.NewPanelComponent(&zgui.PanelOptions{
		Direction: zgui.DirColumn,
	})
	constraints = zgui.DefaultConstraints()
	// constraints.SetHeight(zgui.NewFitConstraint(panel))
	window.Add(panel, constraints)

	constraints = zgui.DefaultConstraints()
	constraints.SetHeight(zgui.NewPixelConstraint(20))
	constraints.SetWidth(zgui.NewPixelConstraint(20))
	checkbox := zgui.NewCheckboxComponent(&zgui.CheckboxOptions{
		Box: &zgui.BoxOptions{
			Color:     rl.Black,
			Roundness: 0.0,
			Segments:  1,
			LineThick: 1,
		},
		Mark: &zgui.TextOptions{
			Color: rl.Black,
		},
	})
	panel.Add(checkbox, constraints)

	constraints = zgui.DefaultConstraints()
	constraints.SetHeight(zgui.NewPixelConstraint(20))
	constraints.SetWidth(zgui.NewRelativeConstraint(0.8))
	slider := zgui.NewSliderComponent(&zgui.SliderOptions{
		Bar: &zgui.BoxOptions{
			Color:     rl.Black,
			Roundness: 0.5,
			Segments:  30,
		},
		Marker: &zgui.BoxOptions{
			Color:     rl.Yellow,
			Roundness: 0.5,
			Segments:  30,
		},
	})
	panel.Add(slider, constraints)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		display.Update(rl.GetFrameTime())
		display.Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
