package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	len, err := f.WriteString("hello, world!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", len)

	arquivo, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	println(string(arquivo))

	arquivo2, err := os.Open("file.txt")
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
