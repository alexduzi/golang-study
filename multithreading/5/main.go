package main

import "fmt"

func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("receiving %d\n", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //indica que nada mais vai se inserido no canal
}
