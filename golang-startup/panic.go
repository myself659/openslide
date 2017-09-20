package main

import (
	"fmt"
)

func safedivide(a, b int) int {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("occur error:", p)
		}
	}()

	if a == 0 {
		panic("invalid num zero")
	}
	fmt.Println("result:", b/a)
	return b / a
}

func main() {
	safedivide(1, 35)
	safedivide(0, 10)
}
