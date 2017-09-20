package main

import (
	"fmt"
	"sync"
)

var count int
var once sync.Once

func allclosure() int {

	onceDo := func() {
		count = 100
	}
	once.Do(onceDo)
	count := count + 1
	return count
}

func onceHead() {
	onceDo := func() {
		fmt.Println("start")
	}
	var once sync.Once
	for i := 0; i < 10; i++ {
		once.Do(onceDo)
		fmt.Println(i)
	}
}

func main() {
	onceHead()
}
