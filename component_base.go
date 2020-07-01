package zgui

import (
	rl "github.com/xzebra/raylib-go/raylib"
)

type baseComponent struct {
	IConstraints
	components []IComponent

	State GuiState
}

func newBaseComponent() *baseComponent {
	return &baseComponent{
		IConstraints: DefaultConstraints(),
		State:        StateNormal,
	}
}

func (b *baseComponent) setConstraints(constraints IConstraints) {
	// b.IConstraints = constraints
	b.IConstraints.SetX(constraints.GetXConstraint())
	b.IConstraints.SetY(constraints.GetYConstraint())
	b.IConstraints.SetWidth(constraints.GetWidthConstraint())
	b.IConstraints.SetHeight(constraints.GetHeightConstraint())
}

func (b *baseComponent) GetConstraints() IConstraints {
	return b.IConstraints
}

func (b *baseComponent) Add(component IComponent, constraints IConstraints) {
	constraints.setParent(b.GetConstraints())
	component.setConstraints(constraints)

	b.components = append(b.components, component)
}

func (b *baseComponent) Update(dt float32) {
	for _, component := range b.components {
		component.Update(dt)
	}
}

func (b *baseComponent) Draw() {
	for _, component := range b.components {
		component.Draw()
	}
}

func (b *baseComponent) TouchInBounds() bool {
	if !rl.IsGestureDetected(rl.GestureTap) {
		return false
	}
	touchPosition := rl.GetTouchPosition(0)
	return b.MouseInBounds(int32(touchPosition.X), int32(touchPosition.Y))
}

func (b *baseComponent) MouseInBounds(mx, my int32) bool {
	tmp := b.GetBounds()
	bounds := tmp.ToInt32()
	return mx > bounds.X && mx < (bounds.X+bounds.Width) &&
		my > bounds.Y && my < (bounds.Y+bounds.Height)
}

func (b *baseComponent) IsSelected() bool {
	return b.State == StatePressed
}

func (b *baseComponent) GetState() GuiState {
	return b.State
}

func (b *baseComponent) SetState(val GuiState) {
	b.State = val
}
