package zgui

type IDrawable interface {
	Draw()
}

type IUpdateable interface {
	Update(dt float32)
}

type IPosition interface {
	GetX() int32
	SetX() int32
	GetY() int32
	SetY() int32
}

type IRectangle interface {
	IPosition

	GetWidth() int32
	SetWidth() int32
	GetHeight() int32
	SetHeight() int32
}

type IConstraints interface {
	SetX(IConstraint constraint)
	SetY(IConstraint constraint)
	SetWidth(IConstraint constraint)
	SetHeight(IConstraint constraint)

	GetX() IConstraint
	GetY() IConstraint
	GetWidth() IConstraint
	GetHeight() IConstraint
}

type IConstraint interface {
	ValueX(IRectangle box) int32
	ValueY(IRectangle box) int32
	ValueWidth(IRectangle box) int32
	ValueHeight(IRectangle box) int32
}

type IComponent interface {
	IDrawable
	IUpdateable
	IRectangle

	// init creates all component parts
	init()

	GetPadding()
	SetPadding()

	SetConstraints(IConstraints constraints)
	GetConstraints() IConstraints

	Add(IComponent component, IConstraints constraints)
}
