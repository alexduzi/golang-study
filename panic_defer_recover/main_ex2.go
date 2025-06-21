package main

import "fmt"

// this code will execute only in the scope of divide function
// if you put this defer anonymous function code inside the main function
// the fmt.Println("Exit...") will not execute
// the panic function will be called at runtime execution
// in other case we can just call panic("some error") by ourselves
func divide(a, b int) (result int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovering from error: (%v)\n", err)
		}
	}()
	result = a / b
	return
}

func main() {
	fmt.Println("Init...")
	result := divide(1, 0)
	fmt.Println(result)

	result = divide(1, 1)
	fmt.Println(result)

	fmt.Println("Exit...")
}
