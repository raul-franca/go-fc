package matematica

// A primeira letra define a acesso a var e func maiuscolo publico minuscolo

var VarPublica string
var varPrivada string

func Soma[T int | float64](a, b T) T {
	return soma2numeros(a, b)
}

//Iniciando com letra minuscola so eh visível apenas no pacote
func soma2numeros[T int | float64](a, b T) T {
	return a + b
}
