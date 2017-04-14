package state

import "testing"

func TestState(t *testing.T) {
	c := Context{}
	c.ChangeState(new(ConcreteStateBegin))
	c.Request("1")
	//out
	//begin state 1

	c.ChangeState(new(ConcreteStateLoading))
	c.Request("2")
	//out
	//loading state 2

	c.ChangeState(new(ConcreteStateEnd))
	c.Request("3")
	//out
	//end state 3
}
