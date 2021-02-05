package zgui

import "math"

// FitConstraint is used to adapt an element to fit the children components.
// Warning: it can lead to recursion loops if children bounds require parent
// bounds.
type FitConstraint struct {
	*baseConstraint
	Parent IParent
}

// NewFitConstraint creates a fitting constraint object.
func NewFitConstraint(parent IParent) *FitConstraint {
	return &FitConstraint{
		baseConstraint: NewBaseConstraint(),
		Parent:         parent,
	}
}

func (c FitConstraint) GetX() (out float32) {
	out = math.MaxFloat32

	// Get minimum X
	for _, children := range c.Parent.Children() {
		x := children.GetX()
		if x < out {
			out = x
		}
	}

	return
}

func (c FitConstraint) GetY() (out float32) {
	out = math.MaxFloat32

	// Get minimum Y
	for _, children := range c.Parent.Children() {
		y := children.GetY()
		if y < out {
			out = y
		}
	}

	return
}

func (c FitConstraint) GetWidth() float32 {
	minX := c.GetX()
	// 0 value could cause errors if children width is not correctly set
	var maxX float32 = 0

	for _, children := range c.Parent.Children() {
		x := children.GetX()
		width := children.GetWidth()

		if x+width > maxX {
			maxX = x + width
		}
	}

	return maxX - minX
}

func (c FitConstraint) GetHeight() float32 {
	minY := c.GetY()
	// 0 value could cause errors if children width is not correctly set
	var maxY float32 = 0

	for _, children := range c.Parent.Children() {
		y := children.GetY()
		width := children.GetWidth()

		if y+width > maxY {
			maxY = y + width
		}
	}

	return maxY - minY
}

// String returns a string representation of the constraint.
func (c FitConstraint) String() string {
	return "FitConstraint"
}
