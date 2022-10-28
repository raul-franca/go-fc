package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//tamanho, err := file.WriteString("Texto para o aquivo escrito com o file.WriteString ")
	tamanho, err := file.Write([]byte("Texto para o aquivo escrito com o file.Write "))
	if err != nil {
		panic(err)
	}
	fmt.Printf("O  tamanho do arquivo é %d bytes\n", tamanho)

	file.Close()

	f, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	println(string(f))

	f2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
