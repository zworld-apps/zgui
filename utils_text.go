package zgui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	defaultFontSize float32 = 32
)

var guiFont rl.Font
var guiFontSize float32 = defaultFontSize

func SetFont(font rl.Font) {
	if font.Texture.ID > 0 {
		guiFont = font
		guiFontSize = float32(font.BaseSize)
	}
}

func LoadDefaultStyle() {
	guiFont = rl.GetFontDefault()
	guiFontSize = float32(guiFont.BaseSize)
}

func getFontSpacing(fontSize float32) float32 {
	return fontSize / defaultFontSize
}

func GetTextWidth(text string, fontSize float32) float32 {
	return rl.MeasureTextEx(guiFont, text,
		fontSize,
		getFontSpacing(fontSize),
	).X
}

func DrawText(text string, bounds rl.Rectangle, alignment Alignment, fontSize float32, color rl.Color) {
	if text == "" {
		return
	}

	// vertical alignment for pixel perfect
	vAlignOffset := func(h float32) float32 {
		return float32(int32(h) % 2)
	}

	position := rl.Vector2{X: bounds.X, Y: bounds.Y}
	textWidth := GetTextWidth(text, fontSize)
	textHeight := fontSize

	switch alignment {
	case AlignStart:
		position.Y += bounds.Height/2 - textHeight/2 + vAlignOffset(bounds.Height)
	case AlignCenter:
		position.X += bounds.Width/2 - textWidth/2
		position.Y += bounds.Height/2 - textHeight/2 + vAlignOffset(bounds.Height)
	case AlignEnd:
		position.X += bounds.Width - textWidth
		position.Y += bounds.Height/2 - textHeight/2 + vAlignOffset(bounds.Height)
	}

	// ensure pixel-perfect coordinates
	position.X = float32(int32(position.X))
	position.Y = float32(int32(position.Y))

	rl.DrawTextEx(guiFont, text, position,
		fontSize, getFontSpacing(fontSize), color)
}
