package adapter

import "fmt"

type ITarget interface {
	TargetMethod()
}

type Adaptee struct {
}

func (*Adaptee) AdapteeMethod() {
	fmt.Println("AdapteeMethod")
}

type Adapter struct {
	Adaptee
}

//TargetMethod a Adaptee AdapteeMethod adapter to target TargetMethod
func (a *Adapter) TargetMethod() {
	a.Adaptee.AdapteeMethod()
}
