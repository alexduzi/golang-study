package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// wg := sync.WaitGroup{}

	ch := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		// wg.Add(1)
		go get(ctx, ch, rand.Int())
	}

	select {
	case <-ctx.Done():
		fmt.Println("done! cancel all goroutines by timeout!")
	default:
		value := <-ch
		if value {
			cancel()
			fmt.Println("canceled by calling directly cancel")
		}
	}
}

func get(ctx context.Context, ch chan bool, num int) {
	time.Sleep(time.Second * time.Duration(num))

	// do some operation

	ch <- true
}
