package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

// $ ab -n 10000 -c 100 http://localhost:3000/
//demostração do problema de raca condition

var visitates uint64 = 0

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		atomic.AddUint64(&visitates, 1)
		//visitates++
		//m.Unlock()
		time.Sleep(time.Millisecond * 2)
		w.Write([]byte(fmt.Sprintf("qtd visitantes: %d", visitates)))
	})
	http.ListenAndServe(":3000", nil)
}
