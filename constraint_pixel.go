package zgui

type PixelConstraint struct {
	val float32
}

func NewPixelConstraint(val float32) *PixelConstraint {
	return &PixelConstraint{val: val}
}

func (c PixelConstraint) ValueX(box IContainer) float32 {
	return box.GetX() + c.val
}

func (c PixelConstraint) ValueY(box IContainer) float32 {
	return box.GetY() + c.val
}

func (c PixelConstraint) ValueWidth(box IContainer) float32 {
	return c.val
}

func (c PixelConstraint) ValueHeight(box IContainer) float32 {
	return c.val
}
