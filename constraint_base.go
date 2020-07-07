package zgui

import "fmt"

type baseConstraint struct {
	parentConstraints IConstraints
}

// NewBaseConstraint creates the base of all constraints and already implements
// the most basic functions. By the way, a baseConstraint by itself does
// absolutely nothing.
func NewBaseConstraint() *baseConstraint {
	return &baseConstraint{}
}

// setParent sets parent constraints. Parent constraints are
// called <self> object and object parent is <parent>
func (c *baseConstraint) setParent(parent IConstraints) {
	c.parentConstraints = parent
}

// self returns the constraints object which the constraint belongs to.
func (c baseConstraint) self() IConstraints {
	return c.parentConstraints
}

// parent returns the constraints object of the parent of the object which the
// constraint belongs to.
func (c baseConstraint) parent() IConstraints {
	return c.self().getParent()
}

// GetX returns always 0. It is gonna be overriden by inheritance.
func (c baseConstraint) GetX() float32 {
	return 0
}

// GetY returns always 0. It is gonna be overriden by inheritance.
func (c baseConstraint) GetY() float32 {
	return 0
}

// GetWidth returns always 0. It is gonna be overriden by inheritance.
func (c baseConstraint) GetWidth() float32 {
	return 0
}

// GetHeight returns always 0. It is gonna be overriden by inheritance.
func (c baseConstraint) GetHeight() float32 {
	return 0
}

func (c *baseConstraint) move(d float32) {
	return
}

// String returns a string representation of the constraint.
func (c baseConstraint) String() string {
	return fmt.Sprintf("baseConstraint{%v}", c.parentConstraints)
}
