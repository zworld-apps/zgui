package zgui

// FillConstraint represents the fixed position and size of
// parent's rectangle. It is used as the default constraint.
type FillConstraint struct {
	*baseConstraint
}

func NewFillConstraint() *FillConstraint {
	return &FillConstraint{
		baseConstraint: newBaseConstraint(),
	}
}
