package zgui

type IDrawable interface {
	Draw()
}

type IUpdateable interface {
	Update(dt float32)
}

type IPosition interface {
	GetX() float32
	GetY() float32
}

type IContainer interface {
	IPosition

	GetWidth() float32
	GetHeight() float32
}

type IConstraint interface {
	IContainer
	setParent(IConstraints)

	self() IConstraints
	parent() IConstraints
}

type IConstraints interface {
	IContainer

	setParent(IConstraints)
	getParent() IConstraints

	SetX(IConstraint)
	SetY(IConstraint)
	SetWidth(IConstraint)
	SetHeight(IConstraint)

	GetBounds() IContainer
}

type IComponent interface {
	IDrawable
	IUpdateable
	IContainer

	// init creates all component parts
	init()

	setConstraints(IConstraints)
	GetConstraints() IConstraints

	Add(IComponent, IConstraints)
}
