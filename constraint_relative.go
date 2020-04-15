package zgui

// RelativeConstraint relative to parent
type RelativeConstraint struct {
	val float32
}

func NewRelativeConstraint(val float32) *RelativeConstraint {
	return &RelativeConstraint{val: val}
}

func (c RelativeConstraint) ValueX(box IContainer) float32 {
	return box.GetX() * c.val
}

func (c RelativeConstraint) ValueY(box IContainer) float32 {
	return box.GetY() * c.val
}

func (c RelativeConstraint) ValueWidth(box IContainer) float32 {
	return box.GetWidth() * c.val
}

func (c RelativeConstraint) ValueHeight(box IContainer) float32 {
	return box.GetHeight() * c.val
}
