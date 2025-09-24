package main

import (
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 4)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 2
	}()

	select {
	case val := <-ch1:
		println("received ch1 ", val)
	case val := <-ch2:
		println("received ch2 ", val)
	case <-time.After(time.Second * 3):
		println("timeout")
	default:
		println("default")
	}

}
