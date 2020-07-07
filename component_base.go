package zgui

import (
	"zgui/events"

	rl "github.com/xzebra/raylib-go/raylib"
)

type baseComponent struct {
	IConstraints
	events.IObserver

	components []IComponent

	// State GuiState
	sm *stateManager

	// Used for dragging calculation
	lastPos   rl.Vector2
	draggable bool
}

func newBaseComponent() *baseComponent {
	b := &baseComponent{
		IConstraints: DefaultConstraints(),
		IObserver:    events.NewObserver(events.Events()),
	}

	b.sm = newStateManager(b)

	return b
}

func (b *baseComponent) setConstraints(constraints IConstraints) {
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
	// Act according to state
	b.sm.Update(dt)

	// Update child components
	for _, component := range b.components {
		component.Update(dt)
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
	return b.GetState() == StatePressed || b.GetState() == StateFocused
}

func (b *baseComponent) SetDraggable(val bool) {
	b.draggable = val
}

func (b *baseComponent) IsDraggable() bool {
	return b.draggable
}

func (b *baseComponent) SetState(state GuiState) {
	b.sm.Change(state)
}

func (b *baseComponent) GetState() GuiState {
	return b.sm.State()
}

func (b *baseComponent) Draw() {
	if b.GetState() == StateHidden {
		return
	}

	for _, component := range b.components {
		component.Draw()
	}
}
