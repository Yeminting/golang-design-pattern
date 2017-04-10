package observer

import "testing"

func TestObserver(t *testing.T) {
	s := SubjectA{}

	oa := new(ObserverA)
	s.Subscribe(oa)

	ob := new(ObserverB)
	s.Subscribe(ob)

	s.SendMsg("hello world!")

	//out
	//observer a receive msg hello world!
	//observer b receive msg hello world!
}
