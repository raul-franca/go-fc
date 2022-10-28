package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := soma(10, 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func soma(a, b int) (int, error) {
	if (a + b) >= 50 {
		return 0, errors.New("é maior que 50")
	}
	return a + b, nil

}
