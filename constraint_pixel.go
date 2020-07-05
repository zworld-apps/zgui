package zgui

// PixelConstraint is used to give an object an absolute position (in pixels)
// inside its parent.
type PixelConstraint struct {
	IConstraint
	val float32
}

// NewPixelConstraint creates a pixel constraint object.
func NewPixelConstraint(val float32) *PixelConstraint {
	return &PixelConstraint{
		IConstraint: NewBaseConstraint(),
		val:         val,
	}
}

// GetX returns the absolute X position using parent's position as origin.
func (c PixelConstraint) GetX() float32 {
	return c.parent().GetX() + c.val
}

// GetX returns the absolute Y position using parent's position as origin.
func (c PixelConstraint) GetY() float32 {
	return c.parent().GetY() + c.val
}

// GetX returns the absolute Width value.
func (c PixelConstraint) GetWidth() float32 {
	return c.val
}

// GetX returns the absolute Height value.
func (c PixelConstraint) GetHeight() float32 {
	return c.val
}

func (c *PixelConstraint) move(d float32) {
	c.val += d
}

// String returns a string representation of the constraint.
func (c PixelConstraint) String() string {
	return "PixelConstraint"
}
