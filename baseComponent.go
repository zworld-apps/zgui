package zgui

import "fmt"

type baseComponent struct {
	IConstraints
	components []IComponent
}

func newBaseComponent() *baseComponent {
	return &baseComponent{}
}

func (b *baseComponent) setConstraints(constraints IConstraints) {
	b.IConstraints = constraints
}

func (b *baseComponent) GetConstraints() IConstraints {
	return b.IConstraints
}

func (b *baseComponent) Add(component IComponent, constraints IConstraints) {
	fmt.Printf("setParent(%+v)\nsetConstraints(%+v)\n", b.GetConstraints(), constraints)
	constraints.setParent(b.GetConstraints())
	component.setConstraints(constraints)
	b.components = append(b.components, component)
}

func (b *baseComponent) init() {}

func (b *baseComponent) Update(dt float32) {
	for _, component := range b.components {
		component.Update(dt)
	}
}

func (b *baseComponent) Draw() {
	for _, component := range b.components {
		component.Draw()
	}
}
