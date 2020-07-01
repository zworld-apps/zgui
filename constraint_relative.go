package zgui

// RelativeConstraint relative to parent
type RelativeConstraint struct {
	IConstraint
	val float32
}

func NewRelativeConstraint(val float32) *RelativeConstraint {
	return &RelativeConstraint{
		IConstraint: newBaseConstraint(),
		val:         val,
	}
}

func (c *RelativeConstraint) SetRelativeValue(val float32) {
	c.val = val
}

func (c RelativeConstraint) GetX() float32 {
	return c.parent().GetWidth() * c.val
}

func (c RelativeConstraint) GetY() float32 {
	return c.parent().GetHeight() * c.val
}

func (c RelativeConstraint) GetWidth() float32 {
	return c.parent().GetWidth() * c.val
}

func (c RelativeConstraint) GetHeight() float32 {
	return c.parent().GetHeight() * c.val
}

func (c RelativeConstraint) String() string {
	return "RelativeConstraint"
}
