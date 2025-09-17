package main

import (
	"fmt"
	"time"
)

func f(x int) {
	var i int
	for i = 0; i < 10; i++ {
		time.Sleep(time.Duration(x) * time.Millisecond)
		fmt.Print(i, " ", x, "\n")
	}
}

func main() {
	go f(30)
	f(10)
}