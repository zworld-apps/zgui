package zgui

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
