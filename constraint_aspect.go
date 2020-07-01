package zgui

// AspectConstraint to keep aspect ratio between the
// opposite bound
type AspectConstraint struct {
	IConstraint
	val float32
}

func NewAspectConstraint(val float32) *AspectConstraint {
	return &AspectConstraint{
		IConstraint: newBaseConstraint(),
		val:         val,
	}
}

func (c AspectConstraint) GetX() float32 {
	return c.self().GetY() * c.val
}

func (c AspectConstraint) GetY() float32 {
	return c.self().GetX() * c.val
}

func (c AspectConstraint) GetWidth() float32 {
	return c.self().GetHeight() * c.val
}

func (c AspectConstraint) GetHeight() float32 {
	return c.self().GetWidth() * c.val
}

func (c AspectConstraint) String() string {
	return "AspectConstraint"
}
