package zgui

import (
	"fmt"

	rl "github.com/xzebra/raylib-go/raylib"
)

type TextOptions struct {
	Color rl.Color
}

func NewTextOptions() *TextOptions {
	return &TextOptions{
		Color: rl.Black,
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
	rl.DrawText(b.Text, int32(b.GetX()), int32(b.GetY()), int32(b.GetHeight()), b.opt.Color)
	b.baseComponent.Draw()
}

func (b *LabelComponent) String() string {
	return fmt.Sprintf("LabelComponent%+v", b.opt)
}
