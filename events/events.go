package events

type EventID int32

const (
	Hover EventID = iota
	Clicked
	Disabled
	Closed
	Moved
)
