package zgui

import (
	rl "github.com/xzebra/raylib-go/raylib"
)

type fullscreenConstraint struct {
	IConstraint
}

func newFullscreenConstraint() fullscreenConstraint {
	return fullscreenConstraint{
		IConstraint: NewBaseConstraint(),
	}
}

func (c fullscreenConstraint) GetX() float32 {
	return 0
}

func (c fullscreenConstraint) GetY() float32 {
	return 0
}

func (c fullscreenConstraint) GetWidth() float32 {
	return float32(rl.GetScreenWidth())
}

func (c fullscreenConstraint) GetHeight() float32 {
	return float32(rl.GetScreenHeight())
}

func (c fullscreenConstraint) String() string {
	return "fullscreenConstraint"
}
