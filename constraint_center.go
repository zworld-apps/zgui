package zgui

// TODO: add support to check self constraints
// can lead to recursive loops
type CenterConstraint struct{}

func NewCenterConstraint() *CenterConstraint {
	return &CenterConstraint{}
}

func (c CenterConstraint) ValueX(box IContainer) float32 {
	return box.GetX()
}

func (c CenterConstraint) ValueY(box IContainer) float32 {
	return box.GetY()
}

func (c CenterConstraint) ValueWidth(box IContainer) float32 {
	return box.GetWidth()
}

func (c CenterConstraint) ValueHeight(box IContainer) float32 {
	return box.GetHeight()
}
