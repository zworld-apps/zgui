package zgui

type CenterConstraint struct {
	IConstraint
}

func NewCenterConstraint() CenterConstraint {
	return CenterConstraint{
		IConstraint: newBaseConstraint(),
	}
}

func (c CenterConstraint) GetX() float32 {
	return (c.parent().GetX() + c.parent().GetWidth()/2) - c.self().GetWidth()/2
}

func (c CenterConstraint) GetY() float32 {
	return (c.parent().GetY() + c.parent().GetHeight()/2) - c.self().GetHeight()/2
}

func (c CenterConstraint) String() string {
	return "CenterConstraint"
}
