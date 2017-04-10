package observer

import "fmt"

type IObserver interface {
	UpdateMsg(string)
}

type ObserverA struct {
}

func (*ObserverA) UpdateMsg(msg string) {
	fmt.Println("observer a receive msg " + msg)
}

type ObserverB struct {
}

func (*ObserverB) UpdateMsg(msg string) {
	fmt.Println("observer b receive msg " + msg)
}

type ISubject interface {
	Subscribe(IObserver)
	Unsubscribe(IObserver)
	SendMsg(string)
}

type SubjectA struct {
	observers []IObserver
}

func (s *SubjectA) Subscribe(o IObserver) {
	if s.observers == nil {
		s.observers = make([]IObserver, 0)
	}
	s.observers = append(s.observers, o)
}

func (s *SubjectA) Unsubscribe(o IObserver) {
	for _, observer := range s.observers {
		if o == observer {

		}
	}
}

func (s *SubjectA) SendMsg(msg string) {
	if len(s.observers) == 0 {
		return
	}

	for _, observer := range s.observers {
		if observer == nil {
			continue
		}

		observer.UpdateMsg(msg)
	}
}
