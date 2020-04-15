package zgui

// Constraints handles all the constraints that will be
// applied to a UI element
type Constraints struct {
	x, y, width, height IConstraint
	container           IContainer
	parent              IConstraints
}

func (c *Constraints) setParent(parent IConstraints) {
	c.parent = parent
}

func (c *Constraints) getParent() IConstraints {
	return c.parent
}

func (c *Constraints) SetX(constraint IConstraint) {
	constraint.setParent(c)
	c.x = constraint
}

func (c *Constraints) SetY(constraint IConstraint) {
	constraint.setParent(c)
	c.y = constraint
}

func (c *Constraints) SetWidth(constraint IConstraint) {
	constraint.setParent(c)
	c.width = constraint
}

func (c *Constraints) SetHeight(constraint IConstraint) {
	constraint.setParent(c)
	c.height = constraint
}

func (c *Constraints) getParentBounds() IContainer {
	if c.parent == nil {
		return nil
	}
	return c.parent.GetBounds()
}

func (c *Constraints) GetX() float32 {
	return c.x.GetX()
}

func (c *Constraints) GetY() float32 {
	return c.y.GetY()
}

func (c *Constraints) GetWidth() float32 {
	return c.width.GetWidth()
}

func (c *Constraints) GetHeight() float32 {
	return c.height.GetHeight()
}

func (c *Constraints) GetBounds() IContainer {
	return &rectangle{
		c.GetX(),
		c.GetY(),
		c.GetWidth(),
		c.GetHeight(),
	}
}

// DefaultConstraints initializes a Constraints object with
// all constraints set to fill parent
func DefaultConstraints() *Constraints {
	return &Constraints{
		x:      FillConstraint{},
		y:      FillConstraint{},
		width:  FillConstraint{},
		height: FillConstraint{},
	}
}
