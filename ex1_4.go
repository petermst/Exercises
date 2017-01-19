package main

import (
	. "fmt"
	"runtime"
	//"time"
)

var i int = 0

func someGoroutine1(c chan int, done chan bool) {
	for j := 0; j < 999999; j++ {
		i = <-c
		i++
		c <- i
	}
	done <- true
}

func someGoroutine2(c chan int, done chan bool) {
	for k := 0; k < 1000000; k++ {
		i = <-c
		i--
		c <- i
	}
	done <- true
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan int, 1)
	done := make(chan bool, 2)
	c <- i

	go someGoroutine1(c, done)
	go someGoroutine2(c, done)

	<-done
	<-done
	i = <-c
	Println("tallet er: ", i)
}
