package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i int = 0

func someGoroutine1() {
	for j := 0; j < 1000000; j++ {
		i++
	}
}

func someGoroutine2() {
	for k := 0; k < 1000000; k++ {
		i--
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go someGoroutine1()
	go someGoroutine2()

	time.Sleep(100 * time.Millisecond)
	Println("tallet er: ", i)
}
