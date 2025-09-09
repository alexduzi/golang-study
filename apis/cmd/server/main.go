package main

import (
	"fmt"

	"github.com/alexduzi/golang-study/apis/configs"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", cfg)
}
