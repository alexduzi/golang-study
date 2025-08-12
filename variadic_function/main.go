package main

import "fmt"

// criação de um tipo
type ID int

func main() {
	salarios := map[string]int{"alex": 17000, "wesley": 10000}
	fmt.Println(salarios["alex"])
	fmt.Println(salarios["wesley"])
	salarios["carol"] = 90000
	fmt.Println(salarios["carol"])
	delete(salarios, "wesley")
	fmt.Printf("%v\n", salarios)
	println(sum(51, 4))

	// closure
	// funcao anonima
	total := func() int {
		return sum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}() // se auto invoca no final

	println(total)
}

func sum(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

// função variática, reticencias nesse caso funciona como um array de args
// parecido com rest operator em js
func sum2(a ...int) (sum int) {
	for _, val := range a {
		sum += val
	}
	return
}
