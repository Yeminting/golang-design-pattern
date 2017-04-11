package proxy

import "testing"

func TestProxy(t *testing.T) {
	var subject ISubject
	subject = new(RealSubject)

	proxy := Proxy{}
	proxy.setSubject(subject)
	proxy.doSomething()

	//out
	//proxy do something here!
	//realsubject do something here!
}
