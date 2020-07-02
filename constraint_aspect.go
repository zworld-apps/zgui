package zgui

// AspectConstraint keeps aspect ratio between the opposite bound.
// ATENTION: It can lead to recursive loops!
type AspectConstraint struct {
	IConstraint
	val float32
}

// NewAspectConstraint creates an aspect constraint object.
func NewAspectConstraint(val float32) *AspectConstraint {
	return &AspectConstraint{
		IConstraint: NewBaseConstraint(),
		val:         val,
	}
}

// GetX returns the X value of the component according
// to its Y value.
func (c AspectConstraint) GetX() float32 {
	return c.self().GetY() * c.val
}

// GetY returns the Y value of the component according
// to its X value.
func (c AspectConstraint) GetY() float32 {
	return c.self().GetX() * c.val
}

// GetWidth returns the Width value of the component according
// to its Height value.
func (c AspectConstraint) GetWidth() float32 {
	return c.self().GetHeight() * c.val
}

// GetHeight returns the Width value of the component according
// to its Width value.
func (c AspectConstraint) GetHeight() float32 {
	return c.self().GetWidth() * c.val
}

// String returns a string representation of the constraint.
func (c AspectConstraint) String() string {
	return "AspectConstraint"
}
