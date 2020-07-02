package zgui

// RelativeConstraint is used to apply a multiplier value to parent's values.
type RelativeConstraint struct {
	IConstraint
	val float32
}

// NewRelativeConstraint creates a relative constraint object.
func NewRelativeConstraint(val float32) *RelativeConstraint {
	return &RelativeConstraint{
		IConstraint: NewBaseConstraint(),
		val:         val,
	}
}

// SetRelativeValue assigns val as the multiplier value.
func (c *RelativeConstraint) SetRelativeValue(val float32) {
	c.val = val
}

// GetX returns parent's X multiplied by the multiplier.
func (c RelativeConstraint) GetX() float32 {
	return c.parent().GetWidth() * c.val
}

// GetY returns parent's Y multiplied by the multiplier.
func (c RelativeConstraint) GetY() float32 {
	return c.parent().GetHeight() * c.val
}

// GetWidth returns parent's Width multiplied by the multiplier.
func (c RelativeConstraint) GetWidth() float32 {
	return c.parent().GetWidth() * c.val
}

// GetHeight returns parent's Height multiplied by the multiplier.
func (c RelativeConstraint) GetHeight() float32 {
	return c.parent().GetHeight() * c.val
}

// String returns a string representation of the constraint.
func (c RelativeConstraint) String() string {
	return "RelativeConstraint"
}
