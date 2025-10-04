package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	var mu = sync.Mutex{}

	// valor compartilhado entre as goroutines que tentam incrementar
	sharedValue := 0

	for i := 0; i < 2; i++ {
		wg.Add(1)

		// multiplas goroutines acessando e modificando o mesmo dado simultaneamente
		// nesse caso o sharedValue será acessado e modificado por diferentes goroutines simultaneamente
		go incrementValueWithMutex(&sharedValue, &mu, &wg)
	}

	// aguardando a conclusão das goroutines
	wg.Wait()

	// exibindo valor final
	fmt.Printf("final sharedValue: %d\n", sharedValue)
}

func incrementValueWithMutex(value *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	// simulando algum processamento antes da modificação do valor
	// isso pode aumentar a chance de race condition
	for i := 0; i < 10000; i++ {
		// utilizando mutex (mutual exclusion) para garantir exclusividade no acesso ao valor
		mu.Lock()
		*value++
		mu.Unlock()
	}
}
