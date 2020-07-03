package zgui

import "fmt"

import "zgui/events"

type buttonCallback func(b *ButtonComponent)

type ButtonOptions struct {
	Box *BoxOptions

	PressCallback buttonCallback
}

type ButtonComponent struct {
	*BoxComponent
	events.IObserver

	opt *ButtonOptions
}

func NewButtonComponent(options *ButtonOptions) *ButtonComponent {
	b := &ButtonComponent{
		BoxComponent: NewBoxComponent(options.Box),
		IObserver:    events.NewObserver(),
		opt:          options,
	}

	b.On(events.Pressed, b.IObserver, func(_ events.EventID) {
		b.opt.PressCallback(b)
	})

	return b
}

func (b *ButtonComponent) String() string {
	return fmt.Sprintf("ButtonComponent")
}
