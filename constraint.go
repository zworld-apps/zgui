package zgui

import "fmt"

type baseConstraint struct {
	parentConstraints IConstraints
}

func newBaseConstraint() *baseConstraint {
	return &baseConstraint{}
}

// setParent sets parent constraints. Parent constraints are
// called <self> object and object parent is <parent>
func (c *baseConstraint) setParent(parent IConstraints) {
	c.parentConstraints = parent
}

// // Pointer function to expose the variable as we need to preserve the functional
// // way of constraints.
// func (c *baseConstraint) ParentConstraints() *IConstraints {
// 	return &c.parentConstraints
// }

func (c baseConstraint) self() IConstraints {
	return c.parentConstraints
}

func (c baseConstraint) parent() IConstraints {
	return c.self().getParent()
}

func (c baseConstraint) GetX() float32 {
	return c.parent().GetX()
}

func (c baseConstraint) GetY() float32 {
	return c.parent().GetY()
}

func (c baseConstraint) GetWidth() float32 {
	return c.parent().GetWidth()
}

func (c baseConstraint) GetHeight() float32 {
	return c.parent().GetHeight()
}

func (c baseConstraint) String() string {
	return fmt.Sprintf("baseConstraint{%v}", c.parentConstraints)
}
