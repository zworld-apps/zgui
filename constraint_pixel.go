package zgui

type PixelConstraint struct {
	IConstraint
	val float32
}

func NewPixelConstraint(val float32) *PixelConstraint {
	return &PixelConstraint{
		IConstraint: newBaseConstraint(),
		val:         val,
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

func (c PixelConstraint) String() string {
	return "PixelConstraint"
}
