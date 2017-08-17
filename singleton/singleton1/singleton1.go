package singleton1

import "fmt"

type singleton struct{}

var std singleton = singleton{}

func (self *singleton) Do() {
	fmt.Println("singleton1 do")
}

func Do() {
	std.Do()
}
