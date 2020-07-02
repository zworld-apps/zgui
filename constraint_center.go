package zgui

// CenterConstraint positions the object in the center of the axis of the
// constraint.
type CenterConstraint struct {
	IConstraint
}

// NewCenterConstraint creates a center constraint object.
func NewCenterConstraint() CenterConstraint {
	return CenterConstraint{
		IConstraint: NewBaseConstraint(),
	}
}

// GetX returns the X value of the component centering it inside parent object.
func (c CenterConstraint) GetX() float32 {
	return (c.parent().GetX() + c.parent().GetWidth()/2) - c.self().GetWidth()/2
}

// GetY returns the Y value of the component centering it inside parent object.
func (c CenterConstraint) GetY() float32 {
	return (c.parent().GetY() + c.parent().GetHeight()/2) - c.self().GetHeight()/2
}

// String returns a string representation of the constraint.
func (c CenterConstraint) String() string {
	return "CenterConstraint"
}
