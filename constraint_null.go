package zgui

// NoConstraint represents the fixed position and size of
// parent's rectangle. It is used as the default constraint.
type NoConstraint struct{}

func (c NoConstraint) ValueX(box IContainer) float32 {
	return box.GetX()
}

func (c NoConstraint) ValueY(box IContainer) float32 {
	return box.GetY()
}

func (c NoConstraint) ValueWidth(box IContainer) float32 {
	return box.GetWidth()
}

func (c NoConstraint) ValueHeight(box IContainer) float32 {
	return box.GetHeight()
}
