package main

import "fmt"

func main() {
	sal := map[string]int{"Raul": 1000, "Joao": 3000, "Maria": 5000}

	fmt.Println(sal["Raul"])
	delete(sal, "Raul")
	sal["franca"] = 2000
	fmt.Println(sal["franca"])

	//sal1 := make(map[string]int)
	//sal2 := map[string]int{}

	for nome, salario := range sal {
		fmt.Printf("O salario de %s é %d \n", nome, salario)
	}
	fmt.Println("Todos os salarios:")
	for _, salario := range sal {
		fmt.Printf("O salario é %d \n", salario)
	}
}
