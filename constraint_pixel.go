package zgui

type PixelConstraint struct {
	*baseConstraint
	val float32
}

func NewPixelConstraint(val float32) *PixelConstraint {
	return &PixelConstraint{
		baseConstraint: newBaseConstraint(),
		val:            val,
	}
}

func (c PixelConstraint) GetX() float32 {
	return c.parent().GetX() + c.val
}

func (c PixelConstraint) GetY() float32 {
	return c.parent().GetY() + c.val
}

func (c PixelConstraint) GetWidth() float32 {
	return c.val
}

func (c PixelConstraint) GetHeight() float32 {
	return c.val
}
