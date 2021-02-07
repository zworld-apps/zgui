package zgui

import (
	"fmt"
	"zgui/events"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WindowOptions struct {
	Bar      *BoxOptions
	CloseBtn *TextOptions
	Content  *BoxOptions

	RequiresFocus bool
	// Undraggable sets window with a fixed position.
	Undraggable bool
}

var DefaultWindowOptions = &WindowOptions{
	Bar: &BoxOptions{
		Color:     rl.Blue,
		Roundness: 0.3,
		Segments:  8,
	},
	CloseBtn: &TextOptions{
		Color: rl.White,
	},
	Content: &BoxOptions{
		Color:     rl.White,
		Roundness: 0,
		Segments:  0,
	},
}

type WindowComponent struct {
	*baseComponent

	ID      string
	clicked bool

	Bar      *BoxComponent
	CloseBtn *LabelComponent
	Content  *BoxComponent

	opt *WindowOptions
}

const windowBarHeight = 20

func NewWindowComponent(options *WindowOptions) *WindowComponent {
	w := &WindowComponent{
		baseComponent: newBaseComponent(),
		Bar:           NewBoxComponent(options.Bar),
		CloseBtn:      NewLabelComponent("X", options.CloseBtn),
		Content:       NewBoxComponent(options.Content),
		opt:           options,
	}

	// Add window bar to base component
	w.baseComponent.Add(w.Bar, &Constraints{
		x:      NewFillConstraint(),
		y:      NewFillConstraint(),
		width:  NewFillConstraint(),
		height: NewPixelConstraint(windowBarHeight),
	})

	// We will be able to drag the bar element
	w.Bar.SetDraggable(!options.Undraggable)
	// After marking it as draggable, we have to handle
	// the dragging event
	w.Bar.On(events.Dragged, func() {
		// Avoid window from leaving display
		mPos := holdInsideWindow(rl.GetMousePosition())
		// Move the whole window component
		w.baseComponent.IConstraints.move(
			mPos.X-w.Bar.lastPos.X,
			mPos.Y-w.Bar.lastPos.Y,
		)
	})

	// Add close button to window bar component
	w.Bar.Add(w.CloseBtn, &Constraints{
		x: NewOperationalConstraint(func(c IConstraint) float32 {
			// Position the X button at the end of the window bar
			return c.parent().GetX() + (c.parent().GetWidth() - windowBarHeight)
		}),
		y:      NewFillConstraint(),
		width:  NewAspectConstraint(1.0),
		height: NewFillConstraint(),
	})

	w.CloseBtn.On(events.Pressed, func() {
		w.Close()
	})

	// Add content box to base component
	w.baseComponent.Add(w.Content, &Constraints{
		x:     NewFillConstraint(),
		y:     NewPixelConstraint(windowBarHeight),
		width: NewFillConstraint(),
		height: NewRelativeConstraint(windowBarHeight).SetOperation(func(x, y float32) float32 {
			return x - y
		}),
	})

	return w
}

func (w *WindowComponent) GetID() string {
	return w.ID
}

func (w *WindowComponent) setID(v string) {
	w.ID = v
}

func (w *WindowComponent) Open() {
	w.SetState(StateNormal)
	w.Notify(events.Opened)
}

func (w *WindowComponent) Close() {
	w.SetState(StateHidden)
	w.Notify(events.Closed)
}

func (w *WindowComponent) IsOpen() bool {
	return w.GetState() != StateHidden
}

func (w *WindowComponent) RequiresFocus() bool {
	return w.opt.RequiresFocus
}

func (w *WindowComponent) Update(dt float32) {
	if w == nil {
		return
	}

	if w.GetState() == StateHidden {
		return
	}
	w.baseComponent.Update(dt)
}

func (w *WindowComponent) Add(component IComponent, constraints IConstraints) {
	w.Content.Add(component, constraints)
}

func (w *WindowComponent) String() string {
	return fmt.Sprintf("WindowComponent%+v", w.opt)
}
