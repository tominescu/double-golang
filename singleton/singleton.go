package main

import (
	"github.com/tominescu/double-golang/singleton/singleton1"
	"github.com/tominescu/double-golang/singleton/singleton2"
)

func main() {
	for i := 0; i < 100; i++ {
		singleton1.Do()
	}
	for i := 0; i < 100; i++ {
		singleton2.GetInstance().Do()
	}
}
