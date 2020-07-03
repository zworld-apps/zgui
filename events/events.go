package events

type EventID int32

const (
	Hovered EventID = iota
	Unhovered
	Pressed
	Released
	Focused
	Unfocused
	Enabled
	Disabled
	Opened
	Closed
	Moved
)
