package facade

import "testing"

func TestFacade(t *testing.T) {
	s1 := new(SubSystem1)
	s2 := new(SubSystem2)
	s3 := new(SubSystem3)

	f := Facade{}
	f.SetSubSystem(s1, s2, s3)
	f.FacadeMethod()

	//subsystem1 method call
	//subsystem2 method call
	//subsystem3 method call
}
