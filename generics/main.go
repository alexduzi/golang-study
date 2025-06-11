package main

import "fmt"

func main() {
	fmt.Println("Iniciando...")

	slice1 := []int{5, 1, 2, 3}
	slice2 := []string{"a", "e", "f", "b"}

	newInts := reverse(slice1)
	newStrings := reverse(slice2)

	fmt.Println(newInts)
	fmt.Println(newStrings)
}

type constraintReverseList interface {
	int | string
}

func reverse[T constraintReverseList](slice []T) []T {
	maxLen := len(slice)

	newInts := make([]T, maxLen)

	newIntsLen := maxLen - 1

	for i := 0; i < maxLen; i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}

	return newInts
}
