package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var matriz [][]string
	var size_of_the_universe int
	var seeds int
	matriz = make([][]string, 1)

	numLinhas := 0
	numColunas := 0
	for {
		//fmt.Println("Size of the universe and seeds: ")
		fmt.Scan(&size_of_the_universe, &seeds)
		if size_of_the_universe > 0 {
			break
		}
	}

	numLinhas = size_of_the_universe
	numColunas = size_of_the_universe
	for i := 0; i < numLinhas; i++ {
		matriz = append(matriz, []string{})
		var coluna []string
		for c := 0; c < numColunas; c++ {
			if rand.Intn(2) == 1 {
				coluna = append(coluna, "O")
			} else {
				coluna = append(coluna, " ")
			}

		}
		matriz[i] = append(matriz[i], coluna...)
	}

	for _, strings := range matriz {
		if len(strings) != 0 {
			for _, s := range strings {
				fmt.Print(s)
			}
			fmt.Print("\n")
		}
	}

}
