package main

import (
	"fmt"
)

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	raul := Cliente{
		Nome:  "Raul",
		Idade: 34,
		Ativo: true,
	}
	raul.Cidade = "Recife"
	fmt.Printf("O cliente %s tem %d mora em %s e esta %t", raul.Nome, raul.Idade, raul.Cidade, raul.Ativo)
}
