package singleton

import "sync"

type Singleton struct {
}

var singleInstance *Singleton
var one sync.Once

//CreateInstance create only one instance
func CreateInstance() *Singleton {
	//make sure create singleInstance once
	one.Do(func() {
		singleInstance = new(Singleton)
	})

	return singleInstance
}
