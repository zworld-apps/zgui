package zgui

import (
	rl "github.com/Lachee/raylib-goplus/raylib"
)

type BoxComponent struct {
	*baseComponent
	color rl.Color
}

func NewBoxComponent(color rl.Color) *BoxComponent {
	return &BoxComponent{
		baseComponent: newBaseComponent(),
		color:         color,
	}
}

func (b *BoxComponent) init() {
}

func (b *BoxComponent) Draw() {
	rl.DrawRectangleRec(
		rl.Rectangle{
			b.GetX(),
			b.GetY(),
			b.GetWidth(),
			b.GetHeight(),
		},
		b.color,
	)
}
