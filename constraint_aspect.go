package zgui

// TODO: AspectConstraint aspect ratio relative to other constraints
type AspectConstraint struct {
	val float32
}

func NewAspectConstraint(val float32) *AspectConstraint {
	return &AspectConstraint{val: val}
}

func (c AspectConstraint) ValueX(box IContainer) float32 {
	return box.GetX() * c.val
}

func (c AspectConstraint) ValueY(box IContainer) float32 {
	return box.GetY() * c.val
}

func (c AspectConstraint) ValueWidth(box IContainer) float32 {
	return box.GetWidth() * c.val
}

func (c AspectConstraint) ValueHeight(box IContainer) float32 {
	return box.GetHeight() * c.val
}
