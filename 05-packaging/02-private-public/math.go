package math02

const ConstPulica = "Tudo que é declarado com letra maiúscula é exportado"

type math struct {
	a int
	b int
	C int
}

func NewMath(a, b int) math {
	return math{a: a, b: b}
}

func (m math) Add() int {
	return m.a + m.b
}
