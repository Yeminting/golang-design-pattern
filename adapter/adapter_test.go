package adapter

import "testing"

func TestAdapter(t *testing.T) {
	var target ITarget
	target = &Adapter{}
	target.TargetMethod()
	//out
	//Adaptee Class Method

	adapter := &AdapterObj{}
	adaptee := new(AdapteeObj)
	adapter.SetAdaptee(adaptee)
	target = adapter
	target.TargetMethod()
	//out
	//Adaptee Object Method
}
