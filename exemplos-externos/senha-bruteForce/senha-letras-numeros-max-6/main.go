package main

import (
	"fmt"
	"time"
)

func main() {
	var targetPassword string // Senha que queremos descobrir
	fmt.Println("Este codigo é demorado.")
	fmt.Print("Digite a senha alfanumérica de até 6 caracteres: ")
	fmt.Scanln(&targetPassword)

	startTime := time.Now() // Início da contagem de tempo
	totalIterations := 0    // Contador de iterações

	for guess := "000000"; guess <= "zzzzzz"; guess = nextPassword(guess) {
		totalIterations++ // Incrementa o contador de iterações

		if guess == targetPassword {
			fmt.Println("Senha encontrada:", guess)
			break
		}
	}

	endTime := time.Now()                           // Fim da contagem de tempo
	elapsedTime := endTime.Sub(startTime).Seconds() // Tempo decorrido em segundos
	totalIterations--                               // Desconta a última iteração (senha encontrada)

	fmt.Println("Tempo decorrido:", elapsedTime, "s")
	fmt.Println("Quantidade de iterações:", totalIterations)
	fmt.Println("Final do programa")
}

func nextPassword(password string) string {
	// Função auxiliar para gerar a próxima senha alfanumérica
	chars := []byte(password)
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] < 'z' {
			chars[i]++
			break
		}
		chars[i] = '0'
	}
	return string(chars)
}
