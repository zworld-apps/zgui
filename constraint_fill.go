package zgui

// FillConstraint represents the fixed position and size of
// parent's rectangle. It is used as the default constraint.
type FillConstraint struct {
	IConstraint
}

func NewFillConstraint() *FillConstraint {
	return &FillConstraint{
		IConstraint: newBaseConstraint(),
	}
}

func (c FillConstraint) String() string {
	return "FillConstraint"
}
