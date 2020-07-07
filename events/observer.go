package events

import (
	"errors"
)

type eventCallback func()

type IObserver interface {
	On(EventID, eventCallback) error
	Notify(EventID) error
}

type ISubject interface {
	Attach(eventCallback)
	Detach()
	Notify()
}

type Subject struct {
	callback eventCallback
}

func NewSubject() *Subject {
	return &Subject{}
}

func (s *Subject) Attach(callback eventCallback) {
	s.callback = callback
}

func (s *Subject) Detach() {
	s.callback = nil
}

func (s *Subject) Notify() {
	if s.callback != nil {
		s.callback()
	}
}

type Observer struct {
	subjects map[EventID]ISubject
}

func NewObserver(observableEvents []EventID) *Observer {
	obs := &Observer{
		subjects: make(map[EventID]ISubject),
	}

	for _, event := range observableEvents {
		obs.subjects[event] = NewSubject()
	}

	return obs
}

func (o *Observer) On(id EventID, callback eventCallback) error {
	if _, exists := o.subjects[id]; !exists {
		return errors.New("event type not supported by component")
	}

	o.subjects[id].Attach(callback)

	return nil
}

func (o *Observer) Notify(id EventID) error {
	if subject, exists := o.subjects[id]; exists {
		subject.Notify()
		return nil
	}

	return errors.New("can't notify unregistered event")
}
