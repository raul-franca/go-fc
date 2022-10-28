package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//log e escrito co stdout do servidor
	log.Print("Request iniciada ")
	defer log.Print("Request finalizada") //defer apenas para ser executado por último
	select {
	case <-time.After(time.Second * 3):
		log.Print("Request processada com sucesso!")
		w.Write([]byte("Request processada com sucesso!"))
	case <-ctx.Done():
		log.Print("Requeste cancelada pelo cliente")
	}
}
