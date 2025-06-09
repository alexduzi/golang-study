package main

import (
	"fmt"
	model "structestudo/model"
	"time"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua:    "Rua X",
		Numero: 616,
		Cidade: "Campinas",
	}

	pessoa := model.Pessoa{
		Nome:             "Alex Duzi",
		Endereco:         endereco,
		DataDeNascimento: time.Date(1991, 02, 01, 0, 0, 0, 0, time.Local),
	}

	fmt.Println(endereco)
	fmt.Println(endereco.Cidade)
	endereco.Numero = 99
	fmt.Println(endereco.Numero)

	fmt.Println(pessoa)

	// idade := model.CalculaIdade(pessoa)
	pessoa.CalculaIdade()

	fmt.Println(pessoa.Idade)

	automovel := model.Automovel{
		Ano:    2022,
		Placa:  "XPTO",
		Modelo: "CG",
	}

	moto := model.Moto{
		Automovel:   automovel,
		Cilindradas: 125,
	}

	fmt.Println(moto)
}
