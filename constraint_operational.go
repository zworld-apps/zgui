package zgui

type operation func(IConstraint) float32

// OperationalConstraint is similar to RelativeConstraint but allows to
// use any parent bound to calculate.
type OperationalConstraint struct {
	IConstraint
	op operation
}

// NewOperationalConstraint creates a operational constraint object.
func NewOperationalConstraint(op operation) *OperationalConstraint {
	return &OperationalConstraint{
		IConstraint: NewBaseConstraint(),
		op:          op,
	}
}

// GetX returns the result of calling the operation.
func (c OperationalConstraint) GetX() float32 {
	return c.op(c)
}

// GetY returns parent's Y multiplied by the multiplier.
func (c OperationalConstraint) GetY() float32 {
	return c.op(c)
}

// GetWidth returns parent's Width multiplied by the multiplier.
func (c OperationalConstraint) GetWidth() float32 {
	return c.op(c)
}

// GetHeight returns parent's Height multiplied by the multiplier.
func (c OperationalConstraint) GetHeight() float32 {
	return c.op(c)
}

// String returns a string representation of the constraint.
func (c OperationalConstraint) String() string {
	return "OperationalConstraint"
}
