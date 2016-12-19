package singleton2

import (
	"fmt"
	"sync"
)

type singleton struct{}

var s *singleton
var once sync.Once

func (self *singleton) Do() {
	fmt.Println("singleton2 do")
}

func GetInstance() *singleton {
	once.Do(func() {
		s = &singleton{}
	})
	return s
}
