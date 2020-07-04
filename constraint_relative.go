package zgui

type relativeOperation func(x float32) float32

// RelativeConstraint is used to apply an operation to parent's values.
type RelativeConstraint struct {
	IConstraint
	op relativeOperation
}

// NewRelativeConstraint creates a relative constraint object.
func NewRelativeConstraint(op relativeOperation) *RelativeConstraint {
	return &RelativeConstraint{
		IConstraint: NewBaseConstraint(),
		op:          op,
	}
}

// GetX returns parent's X multiplied by the multiplier.
func (c RelativeConstraint) GetX() float32 {
	return c.op(c.parent().GetWidth())
}

// GetY returns parent's Y multiplied by the multiplier.
func (c RelativeConstraint) GetY() float32 {
	return c.op(c.parent().GetHeight())
}

// GetWidth returns parent's Width multiplied by the multiplier.
func (c RelativeConstraint) GetWidth() float32 {
	return c.op(c.parent().GetWidth())
}

// GetHeight returns parent's Height multiplied by the multiplier.
func (c RelativeConstraint) GetHeight() float32 {
	return c.op(c.parent().GetHeight())
}

// String returns a string representation of the constraint.
func (c RelativeConstraint) String() string {
	return "RelativeConstraint"
}
