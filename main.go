package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	exit := make(chan bool)

	go func() {
		for i := 0; i <= 20; i++ {
			ch <- i
		}
		exit <- true
	}()

	defer close(ch)
	defer close(exit)
	
	for {
		select {
		case value := <-ch:
			fmt.Printf("Receiving: -> %v\n", value)
		case value := <-exit:
			fmt.Println("End!", value)
		default:
			fmt.Println("no message received")
		}
		time.Sleep(time.Second)
	}
}
