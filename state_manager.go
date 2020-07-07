package zgui

import (
	"fmt"
)

type iStateManager interface {
	Change(id GuiState)
	Update(dt float32)
	State() GuiState
	Component() *baseComponent
}

type stateManager struct {
	component *baseComponent
	Current   GuiState
}

func newStateManager(component *baseComponent) *stateManager {
	return &stateManager{
		component: component,
		Current:   StateNormal,
	}
}

func (sm *stateManager) Change(id GuiState) {
	// check if state changes
	if id == sm.Current {
		return
	}

	// check if state behavior exists
	state, exists := globalStateBehaviors[id]
	if !exists {
		fmt.Printf("zgui: behavior of state %v not declared\n", id)
		return
	}

	// run Exit function of current state
	current := globalStateBehaviors[sm.Current]
	if current.Exit != nil {
		current.Exit(sm)
		// current.Inited = false
	}
	if !current.ExitEvent.IsNull() {
		sm.component.Notify(current.ExitEvent)
	}

	// run Enter function of next state
	if state.Enter != nil {
		state.Enter(sm)
		// state.Inited = true
	}
	if !state.EnterEvent.IsNull() {
		sm.component.Notify(current.EnterEvent)
	}

	// update current state
	sm.Current = id
}

func (sm *stateManager) Update(dt float32) {
	current := globalStateBehaviors[sm.Current]

	// if not yet initialized, run Enter function
	// if !current.Inited {
	// 	current.Enter(sm.component)
	// 	current.Inited = true
	// }

	if current.Update != nil {
		current.Update(sm, dt)
	}
}

func (sm *stateManager) State() GuiState {
	return sm.Current
}

func (sm *stateManager) Component() *baseComponent {
	return sm.component
}
