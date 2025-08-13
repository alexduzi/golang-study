package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int     `json:"numero"`
	Saldo  float64 `json:"saldo"`
}

func main() {
	conta := Conta{123456, 100000.00}
	fmt.Printf("%+v\n", conta)

	contaJson, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(contaJson))

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(conta)
	if err != nil {
		panic(err)
	}

	var contaX Conta
	err = json.Unmarshal(contaJson, &contaX)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal %+v\n", conta)
}
