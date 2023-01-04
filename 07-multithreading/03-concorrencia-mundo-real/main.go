package main

import (
	"fmt"
	"net/http"
	"time"
)

// $ ab -n 10000 -c 100 http://localhost:3000/
//demostração do problema de raca condition

var visitates uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visitates++
		time.Sleep(time.Millisecond * 2)
		w.Write([]byte(fmt.Sprintf("qtd visitantes: %d", visitates)))
	})
	http.ListenAndServe(":3000", nil)
}
