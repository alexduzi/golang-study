package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// go run cep/main.go https://viacep.com.br/ws/01001000/json/
func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Printf("erro ao realizar requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("erro ao ler resposta: %v\n", err)
		}

		var cepInfo ViaCep
		json.Unmarshal(res, &cepInfo)

		file, err := os.Create("./cidade.txt")
		if err != nil {
			fmt.Printf("erro ao criar arquivo cidade.txt: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", cepInfo.Cep, cepInfo.Localidade, cepInfo.Uf))
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Printf("Cidade: %s\n", cepInfo.Localidade)
	}
}
