package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	type Cursos []Curso

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles("content.html"))
		err := t.Execute(writer, Cursos{
			{"Go", 40},
			{"java", 30},
			{"python", 60},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
