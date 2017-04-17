package adapter

import "fmt"

type AdapteeObj struct {
}

func (*AdapteeObj) AdapteeMethod() {
	fmt.Println("Adaptee Object Method")
}

type AdapterObj struct {
	adaptee *AdapteeObj
}

func (a *AdapterObj) SetAdaptee(adaptee *AdapteeObj) {
	a.adaptee = adaptee
}

func (a *AdapterObj) TargetMethod() {
	a.adaptee.AdapteeMethod()
}
