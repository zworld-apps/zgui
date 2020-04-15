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

func (c *Constraints) SetX(constraint IConstraint) {
	c.x = constraint
}

func (c *Constraints) SetY(constraint IConstraint) {
	c.y = constraint
}

func (c *Constraints) SetWidth(constraint IConstraint) {
	c.width = constraint
}

func (c *Constraints) SetHeight(constraint IConstraint) {
	c.height = constraint
}

func (c *Constraints) getParentBounds() IContainer {
	if c.parent == nil {
		return nil
	}
	return c.parent.GetBounds()
}

func (c *Constraints) GetX() float32 {
	return c.x.ValueX(c.getParentBounds())
}

func (c *Constraints) GetY() float32 {
	return c.y.ValueY(c.getParentBounds())
}

func (c *Constraints) GetWidth() float32 {
	return c.width.ValueWidth(c.getParentBounds())
}

func (c *Constraints) GetHeight() float32 {
	return c.height.ValueHeight(c.getParentBounds())
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
// all constraints set to NoConstraint
func DefaultConstraints() *Constraints {
	return &Constraints{
		x:      NoConstraint{},
		y:      NoConstraint{},
		width:  NoConstraint{},
		height: NoConstraint{},
	}
}
