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

type fullscreenConstraint struct{}

func newFullscreenConstraint() *fullscreenConstraint {
	return &fullscreenConstraint{}
}

func (c *fullscreenConstraint) ValueX(box IContainer) float32 {
	return 0
}

func (c *fullscreenConstraint) ValueY(box IContainer) float32 {
	return 0
}

func (c *fullscreenConstraint) ValueWidth(box IContainer) float32 {
	return float32(rl.GetScreenWidth())
}

func (c *fullscreenConstraint) ValueHeight(box IContainer) float32 {
	return float32(rl.GetScreenHeight())
}
