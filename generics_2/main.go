package main

import "fmt"

func main() {
	// interface vazia pode lidar com qualquer tipo
	var x interface{} = 10
	var y interface{} = "Hello, World!"
	showType(x)
	showType(y)

	// type assertion
	var minhaVar interface{} = "Alex Duzi"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	if !ok {
		println("conversão não deu certo")
	} else {
		println(res)
	}
	println()

	m := map[string]int{"Alex": 1000, "Bruno": 40000}
	m2 := map[string]float64{"Alex": 1000.90, "Bruno": 40000.54}
	m3 := map[string]MyNumber{"Alex": 1000, "Bruno": 40000}
	fmt.Printf("O valor da soma é: %v\n", Soma(m))
	fmt.Printf("O valor da soma é: %v\n", Soma(m2))
	fmt.Printf("O valor da soma é: %v\n", Soma(m3))
	println(Compara(10, 10.0))
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, val := range m {
		soma += val
	}
	return soma
}
