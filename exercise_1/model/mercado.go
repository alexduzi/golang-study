package model

import (
	"errors"
	"time"
)

type Mercado struct {
	Mercado    string
	Itens      []ItemMercado
	DataCompra time.Time
}

type ItemMercado struct {
	Nome string
}

func Inicializar(nomeMercado string, dataCompra time.Time, itens []ItemMercado) (*Mercado, error) {

	if nomeMercado == "" {
		return nil, errors.New("nome do mercado é obrigatório")
	}

	if len(itens) == 0 {
		return nil, errors.New("itens de compra estão zerados")
	}

	mercado := Mercado{
		Mercado:    nomeMercado,
		Itens:      itens,
		DataCompra: dataCompra, //adiciona 1 dia
	}

	return &mercado, nil
}
