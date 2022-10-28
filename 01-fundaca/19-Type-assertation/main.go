package main

import "fmt"

func main() {

	var minhaVar interface{} = "Raul França"
	println(minhaVar)
	println(minhaVar.(string))

	minhaVar = "10-closure"
	println("Novo valor para minhaVar", minhaVar.(string))

	res, ok := minhaVar.(int) //se ok true convertido com sucesso
	fmt.Printf("varlor res: %v e o ok %v", res, ok)
}
