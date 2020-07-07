package events

// EventID is the type used to handle events.
type EventID int32

const (
	// null is the zero-value event, it does nothing.
	null EventID = iota
	// Hovered is the event triggered when entering hover state.
	Hovered
	// Unhovered is the event triggered when leaving hover state.
	Unhovered
	// Pressed is the event triggered when entering pressed state.
	Pressed
	// Released is the event triggered when leaving pressed state.
	Released
	// Focused is the event triggered when entering focused state.
	Focused
	// Unfocused is the event triggered when leaving focused state.
	Unfocused
	// Dragged is the event triggered when element is being dragged.
	Dragged
	// Opened is the event triggered when element is opened.
	Opened
	// Closed is the event triggered when element is closed.
	Closed
	// Moved
	eventMax // not exported, used for ranged iteration
)

func (e EventID) IsNull() bool {
	return e == null
}

// Events returns a list of all posible events.
func Events() (events []EventID) {
	for i := null + 1; i < eventMax; i++ {
		events = append(events, i)
	}
	return
}
