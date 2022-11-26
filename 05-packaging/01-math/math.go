package math

type Math struct {
	A int
	B int
}

func (m Math) Add() int {
	var r int
	r = m.A + m.B
	return r
}
