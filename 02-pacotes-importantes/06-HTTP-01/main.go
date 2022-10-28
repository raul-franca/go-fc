package main

import "net/http"

func main() {
	http.HandleFunc("/", ola)
	http.ListenAndServe(":8080", nil)
}

func ola(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
