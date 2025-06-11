package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Inicializando...")

	// for i := 0; i < 10000; i++ {
	// 	go showMessage(strconv.Itoa(i))
	// }

	// time.Sleep(time.Duration(time.Hour.Seconds() * float64(5)))

	// var wg sync.WaitGroup
	// wg.Add(3)
	// go callDatabase(&wg)
	// go callApi(&wg)
	// go processInternal(&wg)
	// wg.Wait()

	// i := 0
	// go ChangeNumber(&i, 5)
	// go ChangeNumber(&i, 10)
	// go ChangeNumber(&i, 20)

	// time.Sleep(time.Second * 2)
	// fmt.Println(i)

	// var m sync.Mutex
	// i := 0
	// for x := 0; x < 10000; x++ {
	// 	go func() {
	// 		m.Lock()
	// 		i++
	// 		m.Unlock()
	// 	}()
	// }

	// time.Sleep(time.Second * 5)
	// fmt.Println(i)

	// channel normal
	// channel := make(chan int)

	// channel inicializado com buffer de 100
	channel := make(chan int, 100)
	go setList(channel)

	for v := range channel {
		fmt.Println("recenbendo: ", v)
		time.Sleep(time.Second)
	}
}

// func showMessage(message string) {
// 	fmt.Println(message)
// }

// func callDatabase(wg *sync.WaitGroup) {
// 	time.Sleep(2 * time.Second)

// 	fmt.Println("Finalizado callDatabase")
// 	wg.Done()
// }

// func callApi(wg *sync.WaitGroup) {
// 	time.Sleep(2 * time.Second)

// 	fmt.Println("Finalizado callApi")
// 	wg.Done()
// }

// func processInternal(wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second)

// 	fmt.Println("Finalizado processInternal")
// 	wg.Done()
// }

// func ChangeNumber(i *int, newNumber int) {
// 	*i = newNumber
// }

func setList(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("enviando: ", i)
	}
	close(channel)
}
