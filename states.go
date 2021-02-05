package zgui

import (
	"zgui/events"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GuiState int32

const (
	StateNormal GuiState = iota
	StateHover
	StateFocused
	StatePressed
	StateDisabled
	StateHidden
	StateDragging
)

func (s GuiState) String() string {
	switch s {
	case StateNormal:
		return "StateNormal"
	case StateHover:
		return "StateHover"
	case StateFocused:
		return "StateFocused"
	case StatePressed:
		return "StatePressed"
	case StateDisabled:
		return "StateDisabled"
	case StateHidden:
		return "StateHidden"
	case StateDragging:
		return "StateDragging"
	}

	return "nil"
}

// state representation used by the state manager
type state struct {
	EnterEvent, ExitEvent events.EventID

	Enter, Exit func(sm iStateManager)
	Update      func(sm iStateManager, dt float32)
}

var globalStateBehaviors = map[GuiState]*state{
	StateNormal:   {Update: normalUpdate},
	StateHover:    {Update: hoverUpdate, EnterEvent: events.Hovered, ExitEvent: events.Unhovered},
	StateFocused:  {Update: focusedUpdate, EnterEvent: events.Focused, ExitEvent: events.Unfocused},
	StatePressed:  {Update: pressedUpdate, EnterEvent: events.Pressed, ExitEvent: events.Released},
	StateDragging: {Enter: draggingEnter, Update: draggingUpdate, ExitEvent: events.Released},
	StateHidden:   {EnterEvent: events.Closed, ExitEvent: events.Opened},
}

type inputEvents struct{ Hover, Tapped, Held, Touched bool }

func getInputEvents(b *baseComponent) *inputEvents {
	hover := b.MouseInBounds(rl.GetMouseX(), rl.GetMouseY())
	tapped := rl.IsMouseButtonPressed(rl.MouseLeftButton) ||
		rl.IsGestureDetected(rl.GestureTap)
	held := rl.IsMouseButtonDown(rl.MouseLeftButton) ||
		rl.IsGestureDetected(rl.GestureHold)
	touched := (hover && tapped) || b.TouchInBounds()

	return &inputEvents{hover, tapped, held, touched}
}

func normalUpdate(sm iStateManager, dt float32) {
	input := getInputEvents(sm.Component())

	switch {
	case input.Touched: // if object touched
		sm.Change(StatePressed)
	case input.Hover:
		sm.Change(StateHover)
	}
}

func hoverUpdate(sm iStateManager, dt float32) {
	input := getInputEvents(sm.Component())

	switch {
	case !input.Hover:
		sm.Change(StateNormal) // no longer hovering
	case input.Touched: // if object touched
		sm.Change(StatePressed)
	}
}

func focusedUpdate(sm iStateManager, dt float32) {
	input := getInputEvents(sm.Component())

	switch {
	case !input.Hover && input.Tapped: // no longer focused
		sm.Change(StateNormal)
	case input.Touched: // if object touched
		sm.Change(StatePressed)
	}
}

func pressedUpdate(sm iStateManager, dt float32) {
	input := getInputEvents(sm.Component())

	switch {
	case sm.Component().IsDraggable() && input.Held:
		sm.Change(StateDragging)
	case !input.Touched: // no longer pressed
		sm.Change(StateFocused)
	}
}

func draggingEnter(sm iStateManager) {
	// Store the starting mouse position
	sm.Component().lastPos = rl.GetMousePosition()
}

func draggingUpdate(sm iStateManager, dt float32) {
	input := getInputEvents(sm.Component())

	switch {
	case !input.Held: // no longer dragging
		sm.Change(StateFocused)
	default:
		newPos := holdInsideWindow(rl.GetMousePosition())

		if newPos != sm.Component().lastPos {
			sm.Component().Notify(events.Dragged)
			// Update last position for the next calc
			sm.Component().lastPos = newPos
		}
	}
}
