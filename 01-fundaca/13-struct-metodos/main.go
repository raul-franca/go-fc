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

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)

}

func main() {
	raul := Cliente{
		Nome:  "Raul",
		Idade: 34,
		Ativo: true,
	}
	raul.Cidade = "Recife"
	fmt.Printf("O cliente %s tem %d mora em %s e esta %t \n", raul.Nome, raul.Idade, raul.Cidade, raul.Ativo)

	raul.Desativar()
}
