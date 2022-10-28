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
	curso := Curso{"go", 40}

	temp := template.Must(template.New("CursoTemplete").Parse("curso: {{.Nome}}, carga horaria: {{.CargaHoraria}}"))
	//temp := template.New("cursoTemplete")
	//temp, _ = temp.Parse("curso: {{.Nome}}, carga horaria: {{.CargaHoraria}}")
	err := temp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
