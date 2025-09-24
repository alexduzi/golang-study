package main

import "fmt"

func main() {
	ch := make(chan string)
	go recebe("hello", ch)
	ler(ch)
}

func recebe(nome string, hello chan<- string) {
	// estamos inserindo um valor no canal
	// esse operador (chanel <-) é de inserir no canal
	hello <- nome
}

func ler(data <-chan string) {
	// neste caso estamos retirando a informação do canal
	// <- channel
	value := <-data
	fmt.Println(value)
}
