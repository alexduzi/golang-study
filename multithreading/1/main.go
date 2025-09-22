package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(20)

	go task("A", &wg)

	go task("B", &wg)

	wg.Wait()
}

func task(name string, wg *sync.WaitGroup) {

	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second + 1)
		wg.Done()
	}
}
