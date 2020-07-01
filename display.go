package zgui

import (
	rl "github.com/xzebra/raylib-go/raylib"
)

// ScreenComponent is the component that fits the whole screen
type ScreenComponent struct {
	*baseComponent
}

func GetDisplay() *ScreenComponent {
	screen := &ScreenComponent{
		baseComponent: newBaseComponent(),
	}

	screen.setConstraints(&Constraints{
		x:      newFullscreenConstraint(),
		y:      newFullscreenConstraint(),
		width:  newFullscreenConstraint(),
		height: newFullscreenConstraint(),
	})

	return screen
}

type fullscreenConstraint struct {
	IConstraint
}

func newFullscreenConstraint() fullscreenConstraint {
	return fullscreenConstraint{
		IConstraint: newBaseConstraint(),
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
