package main

import "fmt"

func main() {
	x := 5
	y := &x
	*y = 10

	ImprimirValores(&x, y)

	fmt.Println(x, *y)
	fmt.Println(&x, y)
}

func ImprimirValores(x *int, y *int) {
	*x = 25
}
