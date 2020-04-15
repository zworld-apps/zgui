package zgui

// RelativeConstraint relative to parent
type RelativeConstraint struct {
	*baseConstraint
	val float32
}

func NewRelativeConstraint(val float32) *RelativeConstraint {
	return &RelativeConstraint{
		baseConstraint: newBaseConstraint(),
		val:            val,
	}
}

func (c RelativeConstraint) GetX() float32 {
	return c.parent().GetX() * c.val
}

func (c RelativeConstraint) GetY() float32 {
	return c.parent().GetY() * c.val
}

func (c RelativeConstraint) GetWidth() float32 {
	return c.parent().GetWidth() * c.val
}

func (c RelativeConstraint) GetHeight() float32 {
	return c.parent().GetHeight() * c.val
}
