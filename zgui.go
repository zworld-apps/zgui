package zgui

type GuiState int32

const (
	StateNormal GuiState = iota
	StateFocused
	StatePressed
	StateDisabled
)
