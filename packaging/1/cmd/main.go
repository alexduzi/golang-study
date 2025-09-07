package main

import (
	"fmt"

	"github.com/alexduzi/golang-study/packaging/1/math"
)

func main() {
	fmt.Println("Hello")
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
