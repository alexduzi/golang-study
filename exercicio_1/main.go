package main

import (
	"exercicio1/model"
	"fmt"
	"time"
)

func main() {
	itens := []model.ItemMercado{
		{Nome: "Tomate"},
		{Nome: "Arroz"},
		{Nome: "Cebola"},
	}

	itens = append(itens, model.ItemMercado{Nome: "Feijão"})
	itens = append(itens, model.ItemMercado{Nome: "Carne"})
	itens = append(itens, model.ItemMercado{Nome: "Chocolate"})

	dataCompra := time.Now().AddDate(0, 0, 1)

	mercado, err := model.Inicializar("Pão de açúcar", dataCompra, itens)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Compras no mercado: ", mercado.Mercado)
		fmt.Println("Data de compra no supermercado: ", mercado.DataCompra)
		fmt.Println("Itens a serem comprados: ", mercado.Itens)
	}
}
