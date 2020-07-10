package zgui

type relatedOperation func(x, y float32) float32

// RelatedConstraint is used to apply an operation to parent's values.
type RelatedConstraint struct {
	*RelativeConstraint
	related IConstraints
}

// NewRelatedConstraint creates a related constraint object.
func NewRelatedConstraint(related IConstraints, n float32) *RelatedConstraint {
	r := &RelatedConstraint{
		RelativeConstraint: NewRelativeConstraint(n),
		related:            related,
	}

	r.SetOperation(func(x, y float32) float32 {
		return x + y
	})

	return r
}

func (c RelatedConstraint) GetX() float32 {
	return c.op(c.related.GetX()+c.related.GetWidth(), c.constant)
}

func (c RelatedConstraint) GetY() float32 {
	return c.op(c.related.GetY()+c.related.GetHeight(), c.constant)
}

func (c RelatedConstraint) GetWidth() float32 {
	return c.op(c.related.GetWidth(), c.constant)
}

func (c RelatedConstraint) GetHeight() float32 {
	return c.op(c.related.GetHeight(), c.constant)
}

// String returns a string representation of the constraint.
func (c RelatedConstraint) String() string {
	return "RelatedConstraint"
}
