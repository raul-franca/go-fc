package main

import "fmt"

func main() {

	var x int
	x = 13
	fmt.Println(x)  // valor de x
	fmt.Println(&x) // endereço da x

	var num *int // declarando um ponteiro
	val := new(int)

	num = new(int)
	*num = x // set valor

	// dessa forma val tera o mesmo valor e local na memória que x
	val = &x // set address

	fmt.Println("===pointer num===")
	fmt.Println(*num) // print o valor de x
	fmt.Println(num)  // print address de x
	fmt.Println("===pointer val===")
	fmt.Println(*val) // print a valor de x
	fmt.Println(val)  // print address de x

	//-----resultado ----------
	//	mesmo valor address na memoria diferentes
	//===pointer num===
	//	13
	//0x1400001e0b0
	//===pointer val===
	//	13
	//0x1400001e088

}
