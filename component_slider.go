package zgui

import (
	"fmt"
	"zgui/events"
)

type SliderOptions struct {
	Bar, Marker *BoxOptions
}

type SliderComponent struct {
	*baseComponent

	Bar    *BoxComponent
	Marker *BoxComponent

	opt *SliderOptions
}

func NewSliderComponent(options *SliderOptions) *SliderComponent {
	s := &SliderComponent{
		baseComponent: newBaseComponent(),
		Bar:           NewBoxComponent(options.Bar),
		Marker:        NewBoxComponent(options.Marker),
		opt:           options,
	}

	s.Add(s.Bar, &Constraints{
		x:      NewFillConstraint(),
		y:      NewCenterConstraint(),
		width:  NewFillConstraint(),
		height: NewRelativeConstraint(0.1),
	})

	s.Add(s.Marker, &Constraints{
		x:      NewRelativeConstraint(0),
		y:      NewCenterConstraint(),
		width:  NewRelativeConstraint(0.1),
		height: NewFillConstraint(),
	})

	s.Marker.SetDraggable(true)

	s.Marker.On(events.Dragged, func() {
		mPos := s.GetMouseRelativePos()
		value := clamp(mPos.X, 0, 1)
		s.Marker.GetXConstraint().SetRelativeValue(value)
	})

	return s
}

func (s *SliderComponent) String() string {
	return fmt.Sprintf("SliderComponent%+v", s.opt)
}
