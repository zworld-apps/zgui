package zgui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BoxOptions struct {
	// Roundness is a value between 0.0 and 1.0
	Roundness float32
	// Segments is used to represent the corners
	Segments int32
	Color    rl.Color
	// LineThick is used when the box is lined
	LineThick int32
}

func NewBoxOptions() *BoxOptions {
	return &BoxOptions{
		Roundness: 0.5,
		Segments:  5,
		Color:     rl.Black,
	}
}

type BoxComponent struct {
	*baseComponent
	opt *BoxOptions
}

func NewBoxComponent(options *BoxOptions) *BoxComponent {
	return &BoxComponent{
		baseComponent: newBaseComponent(),
		opt:           options,
	}
}

func (b *BoxComponent) Draw() {
	rl.DrawRectangleRounded(
		rl.Rectangle{
			b.GetX(),
			b.GetY(),
			b.GetWidth(),
			b.GetHeight(),
		},
		b.opt.Roundness,
		b.opt.Segments,
		b.opt.Color,
	)

	b.baseComponent.Draw()
}

func (b *BoxComponent) String() string {
	return fmt.Sprintf("BoxComponent%+v", b.opt)
}

type LineBoxComponent struct {
	*baseComponent
	opt *BoxOptions
}

func NewLineBoxComponent(options *BoxOptions) *LineBoxComponent {
	return &LineBoxComponent{
		baseComponent: newBaseComponent(),
		opt:           options,
	}
}

func (b *LineBoxComponent) Draw() {
	rl.DrawRectangleRoundedLines(
		rl.Rectangle{
			b.GetX(),
			b.GetY(),
			b.GetWidth(),
			b.GetHeight(),
		},
		b.opt.Roundness,
		b.opt.Segments,
		b.opt.LineThick,
		b.opt.Color,
	)

	b.baseComponent.Draw()
}

func (b *LineBoxComponent) String() string {
	return fmt.Sprintf("LineBoxComponent%+v", b.opt)
}
