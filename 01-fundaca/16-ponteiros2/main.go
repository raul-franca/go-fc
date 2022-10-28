package main

import "fmt"

func soma(a, b *int) int {
	*a += 100
	*b += 100
	return *a + *b
}
func somaSemPonteiro(a, b int) int {
	a += 100
	b += 100
	return a + b
}

func main() {

	a := 10
	b := 20

	println("Soma sem o ponteiro:")
	println(somaSemPonteiro(a, b))
	fmt.Printf("a = %d , b = %d \n	", a, b)
	println("Soma com o ponteiro:")
	println(soma(&a, &b))
	fmt.Printf("a = %d , b = %d", a, b)

}
