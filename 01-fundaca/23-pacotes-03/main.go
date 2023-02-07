package main

import (
	"fmt"
	"github.com/raul-franca/go-fc/01-fundaca/21-pacotes-01/matematica"
)

func main() {
	s := matematica.Soma(10, 10)
	fmt.Printf("A soma é %d", s)

	//nao funciona pois soma2numeros nao pode ser acessa direto
	//s := matematica.soma2numeros(10, 10)

}
