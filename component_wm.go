package zgui

import (
	"container/list"
)

const (
	padding    = 5
	paddingBig = 10
)

type WindowManager struct {
	windows        *list.List
	selectedWindow IWindow

	cooldown float32
}

const selectCooldown = 0.2

var Hud = MakeWindowManager()

func MakeWindowManager() *WindowManager {
	return &WindowManager{
		windows: list.New(),
	}
}

func (h *WindowManager) Clear() {
	// FIXME: might not be the best way to clear it
	h.windows = list.New()
	h.selectedWindow = nil
}

func (h *WindowManager) findWindow(id string) *list.Element {
	for e := h.windows.Front(); e != nil; e = e.Next() {
		window, _ := e.Value.(IWindow)
		if window.GetID() == id {
			return e
		}
	}

	return nil
}

func (h *WindowManager) get(id string) (win IWindow, ok bool) {
	// Find window in list
	elem := h.findWindow(id)
	if elem == nil {
		return nil, false
	}

	// Try to get the window element
	win, ok = elem.Value.(IWindow)
	return
}

func (h *WindowManager) AddWindow(id string, win IWindow) {
	// Don't add window if already exists.
	if _, exists := h.get(id); exists {
		return
	}

	win.setID(id)
	h.windows.PushBack(win)
}

func (h *WindowManager) RemoveWindow(id string) {
	elem := h.findWindow(id)
	h.windows.Remove(elem)
}

func (h *WindowManager) IsFocused() bool {
	for e := h.windows.Front(); e != nil; e = e.Next() {
		window, _ := e.Value.(IWindow)
		if window.IsSelected() && window.RequiresFocus() {
			return true
		}
	}
	return false
}

func (h *WindowManager) IsSelected() bool {
	if h.cooldown > 0 {
		return true
	}

	for e := h.windows.Front(); e != nil; e = e.Next() {
		window, _ := e.Value.(IWindow)
		if window.IsSelected() {
			return true
		}
	}
	return false
}

func (h *WindowManager) GetWindow(id string) IWindow {
	win, ok := h.get(id)
	if ok {
		return win
	} else {
		return nil
	}
}

func (h *WindowManager) Update(dt float32) {
	h.cooldown -= dt
	if h.cooldown < 0 {
		h.cooldown = 0
	}

	if h.selectedWindow != nil {
		h.selectedWindow.Update(dt)
		// if a window has ordered to clear hud or deselect itself in
		// its update function, selected window will be nil
		if h.selectedWindow != nil {
			if !h.selectedWindow.IsOpen() {
				h.cooldown = selectCooldown
				h.selectedWindow.SetSelected(false)
				h.selectedWindow = nil
			} else if !h.selectedWindow.IsSelected() {
				h.selectedWindow.SetSelected(false)
				h.selectedWindow = nil
			} else {
				return
			}
		}
	}

	// iterate from the tail as the windows shown in the front are the last
	for e := h.windows.Back(); e != nil; e = e.Prev() {
		if h.selectedWindow != nil {
			return
		}

		window, _ := e.Value.(IWindow)
		window.Update(dt)

		if window.GetState() == StatePressed || window.GetState() == StateDragging {
			if h.selectedWindow != nil && h.selectedWindow != window {
				h.selectedWindow.SetSelected(false)
			}
			h.selectedWindow = window
		}
	}
}

func (h *WindowManager) Draw() {
	for e := h.windows.Back(); e != nil; e = e.Prev() {
		window, _ := e.Value.(IWindow)
		window.Draw()
	}
}

func (h *WindowManager) ErrorWindow(id, error string, closeCallback func()) {
	alertWindow := NewMessage("Error:<br><br>"+error, 24, DefaultWindowOptions)
	// TODO: alertWindow.SetCloseCallback(closeCallback)
	h.AddWindow(id, alertWindow)
	alertWindow.Open()
}
