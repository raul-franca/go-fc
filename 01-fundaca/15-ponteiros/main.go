package main

func main() {
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a // & endereço na memória
	println(b)
	*b = 30 // * valor na memoria
	println(a)
}
