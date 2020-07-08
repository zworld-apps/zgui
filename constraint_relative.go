package zgui

type relativeOperation func(x, y float32) float32

// RelativeConstraint is used to apply an operation to parent's values.
type RelativeConstraint struct {
	IConstraint
	constant float32
	op       relativeOperation
}

// NewRelativeConstraint creates a relative constraint object.
func NewRelativeConstraint(n float32) *RelativeConstraint {
	return &RelativeConstraint{
		IConstraint: NewBaseConstraint(),
		constant:    n,
		op: func(x, y float32) float32 {
			return x * y
		},
	}
}

func (c *RelativeConstraint) SetRelativeValue(val float32) {
	c.constant = val
}

func (c *RelativeConstraint) SetOperation(op relativeOperation) *RelativeConstraint {
	c.op = op
	return c
}

func relativeOutOfBounds(n, objectBound, parentN, parentBound float32) float32 {
	if n+objectBound > parentN+parentBound {
		return n - (n + objectBound - (parentN + parentBound))
	}
	return n
}

// GetX returns parent's X multiplied by the multiplier.
func (c RelativeConstraint) GetX() float32 {
	parentX := c.parent().GetX()
	parentWidth := c.parent().GetWidth()
	width := c.self().GetWidth()

	x := parentX + c.op(parentWidth, c.constant)

	return relativeOutOfBounds(x, width, parentX, parentWidth)
}

// GetY returns parent's Y multiplied by the multiplier.
func (c RelativeConstraint) GetY() float32 {
	parentY := c.parent().GetY()
	parentHeight := c.parent().GetHeight()
	height := c.self().GetHeight()

	y := parentY + c.op(parentHeight, c.constant)

	return relativeOutOfBounds(y, height, parentY, parentHeight)
}

// GetWidth returns parent's Width multiplied by the multiplier.
func (c RelativeConstraint) GetWidth() float32 {
	return c.op(c.parent().GetWidth(), c.constant)
}

// GetHeight returns parent's Height multiplied by the multiplier.
func (c RelativeConstraint) GetHeight() float32 {
	return c.op(c.parent().GetHeight(), c.constant)
}

// String returns a string representation of the constraint.
func (c RelativeConstraint) String() string {
	return "RelativeConstraint"
}
