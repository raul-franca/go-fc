package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://transform.tools/json-to-go
type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	var cep string

	for _, c := range os.Args[1:] {
		cep += c
	}
	//requisicao para o viacep  viacep.com.br/ws/01001000/json/
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Falha na requisição, error: %s \n", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Falha ao ler a resposta, error: %s \n", err)
	}
	fmt.Println(string(res))

	var data ViaCep
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Falha ao fazer o parce da resposta para data, error: %s \n", err)
	}

	file, err := os.Create("cep.txt")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Falha ao criar o arquivo, error: %s \n", err)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("CEP: %s, logradouro: %s, bairro %s, localidade: %s, UF: %s, DDD: %s", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf, data.Ddd))
	fmt.Print("Arquivo cep.txt criado com sucesso!")
}
