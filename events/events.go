package events

type EventID int32

const (
	Null EventID = iota
	Hovered
	Unhovered
	Pressed
	Released
	Focused
	Unfocused
	Enabled
	Disabled
	Dragged
	Opened
	Closed
	Moved
	eventMax // not exported, used for ranged iteration
)

func Events() (events []EventID) {
	for i := Null + 1; i < eventMax; i++ {
		events = append(events, i)
	}
	return
}
