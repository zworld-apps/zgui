package zgui

type CenterConstraint struct {
	*baseConstraint
}

func NewCenterConstraint() *CenterConstraint {
	return &CenterConstraint{
		baseConstraint: newBaseConstraint(),
	}
}

func (c CenterConstraint) GetX() float32 {
	return (c.parent().GetX() + c.parent().GetWidth()/2) - c.self().GetWidth()/2
}

func (c CenterConstraint) GetY() float32 {
	return (c.parent().GetY() + c.parent().GetHeight()/2) - c.self().GetHeight()/2
}
