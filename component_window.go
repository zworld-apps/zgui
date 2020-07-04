package zgui

import "fmt"

type WindowOptions struct {
	Bar     *BoxOptions
	Close   *TextOptions
	Content *BoxOptions
}

type WindowComponent struct {
	*baseComponent

	Bar     *BoxComponent
	Close   *LabelComponent
	Content *BoxComponent

	opt *WindowOptions
}

const windowBarHeight = 20

func NewWindowComponent(options *WindowOptions) *WindowComponent {
	w := &WindowComponent{
		baseComponent: newBaseComponent(),
		Bar:           NewBoxComponent(options.Bar),
		Close:         NewLabelComponent("X", options.Close),
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

	// Add close button to window bar component
	w.Bar.Add(w.Close, &Constraints{
		x: NewOperationalConstraint(func(c IConstraint) float32 {
			// Position the X button at the end of the window bar
			return c.parent().GetX() + (c.parent().GetWidth() - windowBarHeight)
		}),
		y:      NewFillConstraint(),
		width:  NewFillConstraint(),
		height: NewFillConstraint(),
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

func (w *WindowComponent) String() string {
	return fmt.Sprintf("WindowComponent%+v", w.opt)
}
