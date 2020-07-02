package zgui

// screenComponent is the component that fits the whole screen
type screenComponent struct {
	*baseComponent
}

func newScreenComponent() *screenComponent {
	screen := &screenComponent{
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

// screen is a global screen representation.
var screen = newScreenComponent()

// GetDisplay returns the GUI component that represents the whole screen.
func GetDisplay() *screenComponent {
	return screen
}
