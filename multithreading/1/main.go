package main

import (
	"fmt"
	"time"
)

func main() {
	go task("A")
	go task("B")

	time.Sleep(time.Second * 25)
}

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second + 1)
	}
}
