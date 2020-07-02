package zgui

import (
	"zworld/events"

	rl "github.com/xzebra/raylib-go/raylib"
)

type baseComponent struct {
	IConstraints
	events.IObservable

	components []IComponent

	State GuiState
}

func newBaseComponent() *baseComponent {
	return &baseComponent{
		IConstraints: DefaultConstraints(),
		IObservable: NewObservable([]events.EventID{
			events.Hover, events.Clicked,
		}),
		State: StateNormal,
	}
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
	// Check possible events
	hover := tf.MouseInBounds(rl.GetMouseX(), rl.GetMouseY())
	tapped := rl.IsMouseButtonPressed(rl.MouseLeftButton) ||
		rl.IsGestureDetected(rl.GestureTap)
	touched := (hover && tapped) || tf.TouchInBounds()

	if hover {
		b.Notify(events.Hover)
	}

	if tapped || touched {
		b.Notify(events.Clicked)
	}

	// Update child components
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
