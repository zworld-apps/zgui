package zgui

// FillConstraint fits the parent object bounds. It is used as the default
// constraint.
type FillConstraint struct {
	IConstraint
}

// NewFillConstraint creates a fill constraint object.
func NewFillConstraint() *FillConstraint {
	return &FillConstraint{
		IConstraint: NewBaseConstraint(),
	}
}

// GetX returns the X value of the parent object.
func (c FillConstraint) GetX() float32 {
	return c.parent().GetX()
}

// GetY returns the Y value of the parent object.
func (c FillConstraint) GetY() float32 {
	return c.parent().GetY()
}

// GetWidth returns the Width value of the parent object.
func (c FillConstraint) GetWidth() float32 {
	return c.parent().GetWidth()
}

// GetHeight returns the Height value of the parent object.
func (c FillConstraint) GetHeight() float32 {
	return c.parent().GetHeight()
}

// String returns a string representation of the constraint.
func (c FillConstraint) String() string {
	return "FillConstraint"
}
