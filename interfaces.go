package zgui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	move(d float32)
	SetRelativeValue(val float32)
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
	GetParentBounds() rl.Rectangle

	move(dx, dy float32)
}

type IComponent interface {
	IDrawable
	IUpdateable
	IConstraints

	fmt.Stringer

	setConstraints(IConstraints)
	GetConstraints() IConstraints

	GetState() GuiState

	Add(IComponent, IConstraints)

	TouchInBounds() bool
	MouseInBounds(mx int32, my int32) bool
}

type IParent interface {
	Children() []IContainer
}

type IWindow interface {
	IComponent

	GetID() string
	RequiresFocus() bool

	IsOpen() bool
	IsSelected() bool

	SetSelected(v bool)
	setID(id string)
}
