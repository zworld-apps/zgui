package zgui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextOptions struct {
	Color rl.Color
	Align Alignment
}

func NewTextOptions() *TextOptions {
	return &TextOptions{
		Color: rl.Black,
		Align: AlignStart,
	}
}

type LabelComponent struct {
	*baseComponent
	Text string
	opt  *TextOptions
}

func NewLabelComponent(text string, options *TextOptions) *LabelComponent {
	return &LabelComponent{
		baseComponent: newBaseComponent(),
		Text:          text,
		opt:           options,
	}
}

func (b *LabelComponent) Draw() {
	DrawText(b.Text, b.GetBounds(), b.opt.Align, b.GetHeight(), b.opt.Color)
	b.baseComponent.Draw()
}

func (b *LabelComponent) String() string {
	return fmt.Sprintf("LabelComponent%+v", b.opt)
}
