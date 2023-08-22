package main

import (
	"flag"
	"fmt"
)

// go run main.go -message="my message" -secret="terces os"
func main() {
	var inputVar string
	var secretVar string
	flag.StringVar(&inputVar, "message", "", "mensagem para criar um resumo")
	flag.StringVar(&secretVar, "secret", "", "segredo para o resumo")
	flag.Parse()
	fmt.Printf("Calculando hash para: %q\nsecret: %q\n", inputVar, secretVar)

}
