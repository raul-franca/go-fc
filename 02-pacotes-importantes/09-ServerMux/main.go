package main

import "net/http"

type blog struct {
	title string
}

func (b blog) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(b.title))

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		println("error 1, ", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
