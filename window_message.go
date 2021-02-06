package zgui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MessageWindow struct {
	*WindowComponent
}

func NewMessage(text string, fontSize int32, opts *WindowOptions) *MessageWindow {
	msg := &MessageWindow{
		WindowComponent: NewWindowComponent(opts),
	}

	constraints := DefaultConstraints()
	constraints.SetHeight(
		NewRelatedConstraint(
			GetDisplay().GetConstraints(),
			1/3,
		),
	)

	msg.Add(NewLabelComponent(text, &TextOptions{
		Color: rl.Black,
		Align: AlignCenter,
	}), constraints)

	return msg
}
