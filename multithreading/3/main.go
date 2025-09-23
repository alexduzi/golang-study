package main

import "fmt"

func main() {
	canal := make(chan string) // neste momento o canal está vazio

	go func() {
		canal <- "testando canal" // inserindo no canal, canal está cheio
	}()

	msg := <-canal // canal esvazia

	fmt.Println(msg)
}
