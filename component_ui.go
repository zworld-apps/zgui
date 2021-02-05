package zgui

import (
	"container/list"
)

const (
	padding    = 5
	paddingBig = 10
)

type HUD struct {
	windows        list.List
	selectedWindow IWindow

	cooldown float32
}

const selectCooldown = 0.2

var Hud = MakeHUD()

func MakeHUD() *HUD {
	return &HUD{
		windows: list.New(),
	}
}

func (h *HUD) Clear() {
	// FIXME: might not be the best way to clear it
	h.windows = list.New()
	h.selectedWindow = nil
}

func (h *HUD) get(id string) IWindow {
	for e := h.windows.Front(); e != nil; e = e.Next() {
		window, _ := e.(IWindow)
		if window.GetID() == id {
			return window
		}
	}
	return nil
}

func (h *HUD) AddWindow(id string, win IWindow) {
	if h.get(id) {
		return
	}

	win.setID(id)
	h.windows.PushBack(win)
}

func (h *HUD) RemoveWindow(id string) {
	if window, exists := h.get(id); exists {
		h.Remove(window)
	}
}

func (h *HUD) IsFocused() bool {
	for e := h.windows.Front(); e != nil; e = e.Next() {
		window, _ := e.(IWindow)
		if window.IsSelected() && window.RequiresFocus() {
			return true
		}
	}
	return false
}

func (h *HUD) IsSelected() bool {
	if h.cooldown > 0 {
		return true
	}

	for e := h.windows.Front(); e != nil; e = e.Next() {
		window, _ := e.(IWindow)
		if window.IsSelected() {
			return true
		}
	}
	return false
}

func (h *HUD) GetWindow(id string) IWindow {
	e := h.get(id)
	if win != nil {
		return win.(IWindow)
	} else {
		return nil
	}
}

func (h *HUD) Update(dt float32) {
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

		window, _ := e.(IWindow)
		window.Update(dt)

		if window.WasClicked() || window.IsDragged() {
			if h.selectedWindow != nil && h.selectedWindow != window {
				h.selectedWindow.SetSelected(false)
			}
			h.selectedWindow = window
		}
	}
}

func (h *HUD) Draw() {
	for e := h.windows.Back(); e != nil; e = e.Prev() {
		window, _ := e.(IWindow)
		window.Draw()
	})
}

func (h *HUD) ErrorWindow(id, error string, closeCallback func()) {
	alertWindow := NewMessage("Error:<br><br>"+error, 24)
	alertWindow.SetCloseCallback(closeCallback)
	h.AddWindow(id, alertWindow)
	alertWindow.Open()
}
