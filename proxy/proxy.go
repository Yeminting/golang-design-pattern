package proxy

import "fmt"

type ISubject interface {
	doSomething()
}

type RealSubject struct {
}

func (*RealSubject) doSomething() {
	fmt.Println("realsubject do something here!")
}

type Proxy struct {
	realSubject ISubject
}

func (p *Proxy) setSubject(s ISubject) {
	p.realSubject = s
}

//doSomething implement ISubject
func (p *Proxy) doSomething() {
	//do something before or after call realsubject
	fmt.Println("proxy do something here!")
	p.realSubject.doSomething()
}
