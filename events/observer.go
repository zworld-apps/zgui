package events

import (
	"container/list"
	"errors"
)

type eventCallback func(EventID)

type IObservable interface {
	On(EventID, IObserver, eventCallback) error
}

type IObserver interface {
	setCallback(EventID, eventCallback)
	Update()
}

type ISubject interface {
	Attach(IObserver, eventCallback)
	Detach(IObserver)
	Notify()

	GetID() EventID
}

type Observer struct {
	callbacks map[EventID]eventCallback
}

func (o *Observer) Update(id EventID) {
	o.callbacks[id](id)
}

func (o *Observer) setCallback(id EventID, callback eventCallback) {
	o.callbacks[id] = callback
}

type Subject struct {
	id        EventID
	observers *list.List
}

func NewSubject(id EventID) *Subject {
	return &Subject{
		id:        id,
		observers: new(list.List),
	}
}

func (s *Subject) GetID() EventID {
	return s.id
}

func (s *Subject) Attach(obs IObserver, callback eventCallback) {
	obs.setCallback(s.id, callback)
	s.observers.PushBack(obs)
}

func (s *Subject) Detach(observer IObserver) {
	observer.setCallback(s.id, nil)
	for obs := s.observers.Front(); obs != nil; obs = obs.Next() {
		if obs.Value.(IObserver) == observer {
			s.observers.Remove(obs)
		}
	}
}

func (s *Subject) Notify() {
	for obs := s.observers.Front(); obs != nil; obs = obs.Next() {
		observer := obs.Value.(IObserver)
		observer.Update()
	}
}

type Observable struct {
	subjects map[EventID]ISubject
}

func NewObservable(observableEvents []EventID) *Observable {
	obs := &Observable{
		subjects: make(map[EventID]ISubject),
	}

	for _, event := range observableEvents {
		subjects[event] = NewSubject(event)
	}

	return obs
}

func (o *Observable) On(id EventID, observer IObserver, callback eventCallback) error {
	if _, exists := o.subjects[id]; !exists {
		return errors.New("event type not supported by component")
	}

	o.subjects[id].Attach(observer, callback)

	return nil
}

func (o *Observable) Notify(id EventID) {
	o.subjects[id].Notify()
}
