package chain_of_responsibility

import "testing"

func TestResponsibility(t *testing.T) {
	a := new(ConcreteHandlerA)
	b := new(ConcreteHandlerB)
	a.SetHander(b)
	c := new(ConcreteHandlerC)
	b.SetHander(c)

	request := [3]int{5, 50, 500}
	for _, r := range request {
		a.HanderRequest(r)
	}

	//out
	//ConcreteHandlerA handler request 5
	//ConcreteHandlerB handler request 50
	//ConcreteHandlerC handler request 500
}
