package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	color.Blue("Inicializando...")

	panicExample()
	// deferExample()

	fmt.Println("Fim")
}

func deferExample() {
	file, err := os.Create("./settings.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write([]byte("teste"))
}

func panicExample() {
	// função anonima
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover")
		}
	}()

	_, err := os.Open("./panic.txt")

	if err != nil {
		panic(err)
	}
}
