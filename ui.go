package zgui

import (
	rl "github.com/Lachee/raylib-goplus/raylib"
)

// Master is the component that fits the whole screen
type ScreenComponent struct {
	*baseComponent
}

func GetDisplay() *ScreenComponent {
	screen := &ScreenComponent{
		baseComponent: newBaseComponent(),
	}

	constraints := DefaultConstraints()
	constraints.SetX(newFullscreenConstraint())
	constraints.SetY(newFullscreenConstraint())
	constraints.SetWidth(newFullscreenConstraint())
	constraints.SetHeight(newFullscreenConstraint())
	screen.setConstraints(constraints)

	return screen
}

type fullscreenConstraint struct {
	*baseConstraint
}

func newFullscreenConstraint() *fullscreenConstraint {
	return &fullscreenConstraint{
		baseConstraint: newBaseConstraint(),
	}
}

func (c *fullscreenConstraint) GetX() float32 {
	return 0
}

func (c *fullscreenConstraint) GetY() float32 {
	return 0
}

func (c *fullscreenConstraint) GetWidth() float32 {
	return float32(rl.GetScreenWidth())
}

func (c *fullscreenConstraint) GetHeight() float32 {
	return float32(rl.GetScreenHeight())
}
