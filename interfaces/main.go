package main

import (
	"errors"
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	area := r.largura * r.altura
	return area
}

type circulo struct {
	radius float64
}

func (c circulo) area() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

func ExibirGeometria(g geometria) {
	fmt.Println(g.area())
}

func main() {
	retangulo := retangulo{
		largura: 1,
		altura:  2,
	}

	circulo := circulo{
		radius: 3,
	}

	ExibirGeometria(retangulo)
	ExibirGeometria(circulo)

	problema := ProblemaDeNetwork{
		rede:     true,
		hardware: false,
	}

	ExibeError(problema)
	ExibeError(errors.New("Teste"))

	// interface vazia pode armazenar qualquer tipo
	var a interface{}
	a = 5
	fmt.Println(a)

	var lista []interface{}
	lista = append(lista, 10)
	lista = append(lista, 7.5)
	lista = append(lista, true)
	lista = append(lista, "teste")

	for _, valor := range lista {
		if v, ok := valor.(string); ok {
			fmt.Println(v + " string")
		} else {
			fmt.Println(valor)
		}
	}

}

// essa struct implementa a inferface error
type ProblemaDeNetwork struct {
	rede     bool
	hardware bool
}

func (p ProblemaDeNetwork) Error() string {
	if p.rede {
		return "Problema de rede"
	} else if p.hardware {
		return "Problema de hardware"
	} else {
		return "outro problema"
	}
}

func ExibeError(err error) {
	fmt.Println(err.Error())
}
