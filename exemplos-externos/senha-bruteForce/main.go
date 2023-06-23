package main

import (
	"fmt"
	"time"
)

func main() {
	var targetPassword string // Senha que queremos descobrir
	fmt.Print("Digite a senha numérica de 4 dígitos: ")
	fmt.Scanln(&targetPassword)

	startTime := time.Now() // Início da contagem de tempo
	totalIterations := 0    // Contador de iterações

	for guess := 0; guess <= 9999; guess++ {
		// Formata o palpite com zeros à esquerda para ter 4 dígitos
		formattedGuess := fmt.Sprintf("%04d", guess)

		totalIterations++ // Incrementa o contador de iterações

		if formattedGuess == targetPassword {
			fmt.Println("Senha encontrada:", formattedGuess)
			break
		}
	}

	endTime := time.Now()                                // Fim da contagem de tempo
	elapsedTime := endTime.Sub(startTime).Milliseconds() // Tempo decorrido em milissegundos

	fmt.Println("Tempo decorrido:", elapsedTime, "ms")
	fmt.Println("Quantidade de iterações:", totalIterations)
	fmt.Println("Final do programa")
}
