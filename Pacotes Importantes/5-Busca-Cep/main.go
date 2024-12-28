package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
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

func main() {
	for _, cep := range os.Args[1:] {
		req, erro := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", erro)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}
		var data ViaCEP
		erro = json.Unmarshal(res, &data)

		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v", erro)
		}

		file, err := os.Create("cidade.txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Bairro, data.Estado))
		fmt.Println("Cep encontrado com sucesso!")
		fmt.Printf("Cidade: %s\n", data.Localidade)
	}
}
