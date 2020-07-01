package zgui

import (
	rl "github.com/xzebra/raylib-go/raylib"
)

type rectangle rl.Rectangle

func (r *rectangle) GetWidth() float32 {
	return r.Width
}

func (r *rectangle) GetHeight() float32 {
	return r.Height
}

func (r *rectangle) GetX() float32 {
	return r.X
}

func (r *rectangle) GetY() float32 {
	return r.Y
}

func (r *rectangle) SetWidth(v float32) {
	r.Width = v
}

func (r *rectangle) SetHeight(v float32) {
	r.Height = v
}

func (r *rectangle) SetX(v float32) {
	r.X = v
}

func (r *rectangle) SetY(v float32) {
	r.Y = v
}

type intRectangle struct {
	X, Y, Width, Height int32
}

func NewIntRectangle(r IContainer) *intRectangle {
	return &intRectangle{
		X:      int32(r.GetX()),
		Y:      int32(r.GetY()),
		Width:  int32(r.GetWidth()),
		Height: int32(r.GetHeight()),
	}
}
