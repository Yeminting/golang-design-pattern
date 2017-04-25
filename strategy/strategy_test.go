package strategy

import "testing"

func TestStrategy(t *testing.T) {
	c := Context{}
	a := new(Strategy1)
	c.SetStrategy(a)
	c.CallStrategy()
	//out
	//strategy1 call

	b := new(Strategy2)
	c.SetStrategy(b)
	c.CallStrategy()
	//out
	//strategy2 call
}
