package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				fmt.Println("recovering panic1")
			}
			if r == "panic2" {
				fmt.Println("recovering panic2")
			}
		}
	}()

	throwPanic("panic1")

	throwPanic("panic2")
}

func throwPanic(msg string) {
	panic(msg)
}
