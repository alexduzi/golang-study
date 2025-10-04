package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	// valor compartilhado entre as goroutines que tentam incrementar
	sharedValue := 0

	for i := 0; i < 2; i++ {
		wg.Add(1)

		// multiplas goroutines acessando e modificando o mesmo dado simultaneamente
		// nesse caso o sharedValue será acessado e modificado por diferentes goroutines simultaneamente
		go incrementValue(&sharedValue, &wg)
	}

	// aguardando a conclusão das goroutines
	wg.Wait()

	// exibindo valor final
	fmt.Printf("final sharedValue: %d\n", sharedValue)
}

func incrementValue(value *int, wg *sync.WaitGroup) {
	defer wg.Done()

	// simulando algum processamento antes da modificação do valor
	// isso pode aumentar a chance de race condition
	for i := 0; i < 10000; i++ {
		*value++
	}
}
