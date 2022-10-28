package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		println("1", err)
	}
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound) // 404 page not found
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!!!!"))
}
