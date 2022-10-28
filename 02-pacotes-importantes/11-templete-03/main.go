package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	type Cursos []Curso

	temp := template.Must(template.New("content.html").ParseFiles("content.html"))
	err := temp.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"java", 30},
		{"python", 60},
	})
	if err != nil {
		panic(err)
	}
}
