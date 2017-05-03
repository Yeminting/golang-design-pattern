package command

import "testing"

func TestCommand(t *testing.T) {
	r := new(Revicer)
	c := new(CreateCommand)
	c.SerRevicer(r)
	invoker := Invoker{}
	invoker.SetCmd(c)

	invoker.ExcuteCmd()
	invoker.CancelCmd()
	//out
	//revicer do something
	//revicer undo something
}
