package main

import (
	"log"
	"net/http"
)

func main() {

	fileSever := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileSever)
	//mux.Handle("/blog", fileSever)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
