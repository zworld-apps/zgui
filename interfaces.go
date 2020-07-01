package zgui

import (
	"fmt"

	rl "github.com/xzebra/raylib-go/raylib"
)

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

	GetXConstraint() IConstraint
	GetYConstraint() IConstraint
	GetHeightConstraint() IConstraint
	GetWidthConstraint() IConstraint

	GetBounds() rl.Rectangle
}

type IComponent interface {
	IDrawable
	IUpdateable
	IContainer

	fmt.Stringer

	setConstraints(IConstraints)
	GetConstraints() IConstraints

	Add(IComponent, IConstraints)

	TouchInBounds() bool
	MouseInBounds(mx int32, my int32) bool
}
