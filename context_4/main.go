package main

import (
	"context"
	"fmt"
)

func main() {
	// contexto com valores anexados ao contexto
	ctx := context.WithValue(context.Background(), "token", "senha")
	bookHotel(ctx)

}

// por convenção sempre passar o contexto como primeiro parametro
func bookHotel(ctx context.Context) {
	token := ctx.Value("token")

	// convertendo por type assertion
	value, ok := token.(string)
	if ok {
		println("é uma string -> " + value)
	} else {
		println("não conseguiu converter")
	}

	// switch case type assertion
	switch v := token.(type) {
	case string:
		fmt.Printf("variavel do contexto é: %T %v\n", v, v)
	case int64:
		fmt.Printf("variavel do contexto é: %T %v\n", v, v)
	case bool:
		fmt.Printf("variavel do contexto é: %T %v\n", v, v)
	case float64:
		fmt.Printf("variavel do contexto é: %T %v\n", v, v)
	default:
		fmt.Println("não foi possivel encontrar tipo")
	}
}
