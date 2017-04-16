package adapter

import "testing"

func TestAdapter(t *testing.T) {
	target := Adapter{}
	target.TargetMethod()
	//out
	//AdapteeMethod
}
