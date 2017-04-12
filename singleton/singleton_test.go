package singleton

import "testing"
import "fmt"

func TestSingle(t *testing.T) {
	instance1 := CreateInstance()
	instance2 := CreateInstance()

	//show pointer of instance1 and instance2
	fmt.Printf("instance1 %p \n", instance1)
	fmt.Printf("instance2 %p \n", instance2)

	if instance1 != instance2 {
		t.Error("not equal")
	}
}
