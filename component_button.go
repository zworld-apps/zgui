package zgui

import "fmt"

type ButtonComponent struct {
	*BoxComponent
}

func NewButtonComponent(options *BoxOptions) *ButtonComponent {
	return &ButtonComponent{
		BoxComponent: NewBoxComponent(options),
	}
}

func (b *ButtonComponent) Update() {

}

func (b *ButtonComponent) Draw() {
	b.baseComponent.Draw()
}

func (b *ButtonComponent) String() string {
	return fmt.Sprintf("ButtonComponent")
}
