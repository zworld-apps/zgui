package zgui

import (
	"zgui/events"
)

type CheckboxOptions struct {
	Box  *BoxOptions
	Mark *TextOptions
}

type CheckboxComponent struct {
	*LineBoxComponent
	Mark *LabelComponent

	opt *CheckboxOptions
}

func NewCheckboxComponent(options *CheckboxOptions) *CheckboxComponent {
	options.Mark.Align = AlignCenter

	cb := &CheckboxComponent{
		LineBoxComponent: NewLineBoxComponent(options.Box),
		Mark:             NewLabelComponent("X", options.Mark),
		opt:              options,
	}

	// Set mark to fill parent
	cb.Add(cb.Mark, &Constraints{
		x:      NewCenterConstraint(),
		y:      NewCenterConstraint(),
		width:  NewFillConstraint(),
		height: NewFillConstraint(),
	})
	// but uncheck it
	cb.SetChecked(false)

	cb.On(events.Pressed, func() {
		cb.Toggle()
	})

	return cb
}

func (cb *CheckboxComponent) Toggle() {
	cb.SetChecked(!cb.IsChecked())
}

func (cb *CheckboxComponent) IsChecked() bool {
	return cb.Mark.GetState() != StateHidden
}

func (cb *CheckboxComponent) SetChecked(val bool) {
	if val {
		cb.Mark.SetState(StateNormal)
	} else {
		cb.Mark.SetState(StateHidden)
	}
}
