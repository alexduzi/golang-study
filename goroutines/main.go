package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Inicializando...")

	// for i := 0; i < 10000; i++ {
	// 	go showMessage(strconv.Itoa(i))
	// }

	// time.Sleep(time.Duration(time.Hour.Seconds() * float64(5)))

	var wg sync.WaitGroup
	wg.Add(3)
	go callDatabase(&wg)
	go callApi(&wg)
	go processInternal(&wg)
	wg.Wait()
}

// func showMessage(message string) {
// 	fmt.Println(message)
// }

func callDatabase(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)

	fmt.Println("Finalizado callDatabase")
	wg.Done()
}

func callApi(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)

	fmt.Println("Finalizado callApi")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)

	fmt.Println("Finalizado processInternal")
	wg.Done()
}
