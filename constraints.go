package zgui

import (
	"fmt"

	rl "github.com/xzebra/raylib-go/raylib"
)

// Constraints handles all the constraints that will be
// applied to a UI element
type Constraints struct {
	x, y, width, height IConstraint
	container           IContainer
	parent              IConstraints
}

func (c *Constraints) setParent(parent IConstraints) {
	c.parent = parent

	c.x.setParent(c)
	c.y.setParent(c)
	c.width.setParent(c)
	c.height.setParent(c)
}

func (c Constraints) getParent() IConstraints {
	return c.parent
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

func (c Constraints) getParentBounds() rl.Rectangle {
	if c.parent == nil {
		return rl.Rectangle{0, 0, 0, 0}
	}
	return c.parent.GetBounds()
}

func (c Constraints) GetX() float32 {
	return c.x.GetX()
}

func (c Constraints) GetY() float32 {
	return c.y.GetY()
}

func (c Constraints) GetWidth() float32 {
	return c.width.GetWidth()
}

func (c Constraints) GetHeight() float32 {
	return c.height.GetHeight()
}

func (c Constraints) GetBounds() rl.Rectangle {
	return rl.Rectangle{
		c.GetX(),
		c.GetY(),
		c.GetWidth(),
		c.GetHeight(),
	}
}

// move requires X and Y constraints to be movable to actually work.
func (c *Constraints) move(dx, dy float32) {
	c.x.move(dx)
	c.y.move(dy)
}

func (c Constraints) GetXConstraint() IConstraint {
	return c.x
}

func (c Constraints) GetYConstraint() IConstraint {
	return c.y
}

func (c Constraints) GetWidthConstraint() IConstraint {
	return c.width
}

func (c Constraints) GetHeightConstraint() IConstraint {
	return c.height
}

func (c Constraints) String() string {
	return fmt.Sprintf("{ X:%s Y:%s Width:%s Height:%s }", c.x, c.y, c.width, c.height)
}

func emptyConstraints() IConstraints {
	return &Constraints{}
}

// DefaultConstraints initializes a IConstraints object with
// all constraints set to fill parent
func DefaultConstraints() IConstraints {
	return &Constraints{
		x:      NewFillConstraint(),
		y:      NewFillConstraint(),
		width:  NewFillConstraint(),
		height: NewFillConstraint(),
	}
}
