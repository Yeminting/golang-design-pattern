package facade

import "fmt"

type SubSystem1 struct {
}

func (*SubSystem1) Method1() {
	fmt.Println("subsystem1 method call")
}

type SubSystem2 struct {
}

func (*SubSystem2) Method2() {
	fmt.Println("subsystem2 method call")
}

type SubSystem3 struct {
}

func (*SubSystem3) Method3() {
	fmt.Println("subsystem3 method call")
}

type Facade struct {
	sys1 *SubSystem1
	sys2 *SubSystem2
	sys3 *SubSystem3
}

func (f *Facade) SetSubSystem(s1 *SubSystem1, s2 *SubSystem2, s3 *SubSystem3) {
	f.sys1 = s1
	f.sys2 = s2
	f.sys3 = s3
}

func (f *Facade) FacadeMethod() {
	f.sys1.Method1()
	f.sys2.Method2()
	f.sys3.Method3()
}
