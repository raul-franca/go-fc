package main

import (
	"net/http"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}
type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
		t = template.Must(t.ParseFiles(templates...))
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
