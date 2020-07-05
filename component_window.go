package zgui

import (
	"fmt"
	"zgui/events"
)

type WindowOptions struct {
	Bar      *BoxOptions
	CloseBtn *TextOptions
	Content  *BoxOptions
}

type WindowComponent struct {
	*baseComponent

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
	w.Add(w.Bar, &Constraints{
		x:      NewFillConstraint(),
		y:      NewFillConstraint(),
		width:  NewFillConstraint(),
		height: NewPixelConstraint(windowBarHeight),
	})

	w.Bar.On(events.Pressed, func() {
		w.SetState(StateDragging)
	})

	w.Bar.On(events.Released, func() {
		w.SetState(StateFocused)
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
	w.Add(w.Content, &Constraints{
		x:     NewFillConstraint(),
		y:     NewPixelConstraint(windowBarHeight),
		width: NewFillConstraint(),
		height: NewRelativeConstraint(func(x float32) float32 {
			return x - windowBarHeight
		}),
	})

	return w
}

func (w *WindowComponent) Open() {
	w.SetState(StateNormal)
	w.Notify(events.Opened)
}

func (w *WindowComponent) Close() {
	w.SetState(StateHidden)
	w.Notify(events.Closed)
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

func (w *WindowComponent) String() string {
	return fmt.Sprintf("WindowComponent%+v", w.opt)
}
