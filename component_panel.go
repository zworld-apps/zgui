package zgui

import (
	"fmt"
)

type PanelOptions struct {
	Direction Direction
}

type PanelComponent struct {
	*baseComponent

	Components []IComponent

	// margin adapts to the number of elements
	margin float32
	nElems float32

	opt *PanelOptions
}

const panelOcupation = 0.9

func NewPanelComponent(options *PanelOptions) *PanelComponent {
	return &PanelComponent{
		baseComponent: newBaseComponent(),
		opt:           options,
	}
}

func (s *PanelComponent) getComponentXConstraint(last, comp IComponent) IConstraint {
	if s.opt.Direction == DirColumn {
		return NewCenterConstraint()
	}

	if last != nil {
		return NewRelatedConstraint(last, s.margin)
	}
	return NewRelativeConstraint(s.margin)
}

func (s *PanelComponent) getComponentYConstraint(last, comp IComponent) IConstraint {
	if s.opt.Direction == DirRow {
		return NewCenterConstraint()
	}

	if last != nil {
		return NewRelatedConstraint(last, s.margin)
	}
	fmt.Println("relative", s.margin, s.GetY())
	return NewRelativeConstraint(s.margin)

}

func (s *PanelComponent) getComponentHeightConstraint(comp IComponent) IConstraint {
	constraint := comp.GetHeightConstraint()

	// First of all, check if height is direction dependant
	if s.opt.Direction == DirRow {
		return constraint
	}

	// Check if it is a fit constraint
	if _, ok := constraint.(*FitConstraint); ok {
		// Adapt height to number of elements and desired ocupation
		return NewRelativeConstraint(panelOcupation / s.nElems)
	}

	return constraint
}

func (s *PanelComponent) getComponentWidthConstraint(comp IComponent) IConstraint {
	constraint := comp.GetWidthConstraint()

	// First of all, check if width is direction dependant
	if s.opt.Direction == DirColumn {
		return constraint
	}

	// Check if it is a fixed size constraint
	if _, ok := constraint.(*PixelConstraint); ok {
		return constraint
	}

	// Adapt height to number of elements and desired ocupation
	return NewRelativeConstraint(panelOcupation / s.nElems)
}

// buildPanel arranges all objects according to the configuration. It is called
// any time the Component list is updated.
func (s *PanelComponent) buildPanel() {
	s.baseComponent.ClearChildren()

	s.nElems = float32(len(s.Components))

	if s.nElems < 1 {
		return
	}

	s.margin = (1.0 - panelOcupation) / (s.nElems + 1)

	var lastComponent IComponent = nil
	for _, component := range s.Components {
		s.baseComponent.Add(component, &Constraints{
			x:      s.getComponentXConstraint(lastComponent, component),
			y:      s.getComponentYConstraint(lastComponent, component),
			width:  s.getComponentWidthConstraint(component),
			height: s.getComponentHeightConstraint(component),
		})

		lastComponent = component
	}
}

func (s *PanelComponent) Add(component IComponent, constraints IConstraints) {
	s.Components = append(s.Components, component)

	constraints.setParent(s.GetConstraints())
	component.setConstraints(constraints)

	s.buildPanel()
}

func (s *PanelComponent) Children() (out []IContainer) {
	out = make([]IContainer, len(s.Components))
	for i, component := range s.Components {
		out[i] = component
	}

	return
}

func (s *PanelComponent) String() string {
	return fmt.Sprintf("PanelComponent%+v", s.opt)
}
