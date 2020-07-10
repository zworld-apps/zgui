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

	opt *PanelOptions
}

const panelOcupation = 0.9

func NewPanelComponent(options *PanelOptions) *PanelComponent {
	return &PanelComponent{
		baseComponent: newBaseComponent(),
		opt:           options,
	}
}

func (s *PanelComponent) getComponentConstraints(last, comp IComponent, nElems float32) IConstraints {
	margin := (1.0 - panelOcupation) / (nElems + 1)

	switch s.opt.Direction {
	case DirRow:
		var xConstraint IConstraint
		if last != nil {
			xConstraint = NewRelatedConstraint(last, margin)
		} else {
			xConstraint = NewRelativeConstraint(margin)
		}

		return &Constraints{
			x:      xConstraint,
			y:      NewCenterConstraint(),
			width:  NewRelativeConstraint(panelOcupation / nElems),
			height: comp.GetHeightConstraint(),
		}
	case DirColumn:
		var yConstraint IConstraint
		if last != nil {
			yConstraint = NewRelatedConstraint(last, margin)
		} else {
			yConstraint = NewRelativeConstraint(margin)
		}

		return &Constraints{
			x:      NewCenterConstraint(),
			y:      yConstraint,
			width:  comp.GetWidthConstraint(),
			height: NewRelativeConstraint(panelOcupation / nElems),
		}
	}

	return nil
}

// buildPanel arranges all objects according to the configuration. It is called
// any time the Component list is updated.
func (s *PanelComponent) buildPanel() {
	s.baseComponent.ClearChildren()

	nElems := float32(len(s.Components))
	if nElems < 1 {
		return
	}

	var lastComponent IComponent = nil
	for _, component := range s.Components {
		s.baseComponent.Add(component, s.getComponentConstraints(
			lastComponent, component, nElems,
		))
		lastComponent = component
	}
}

func (s *PanelComponent) Add(component IComponent, constraints IConstraints) {
	s.Components = append(s.Components, component)
	s.buildPanel()
}

func (s *PanelComponent) String() string {
	return fmt.Sprintf("PanelComponent%+v", s.opt)
}
