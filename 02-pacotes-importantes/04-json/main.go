package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	var contaX Conta

	//jsonPuro := []byte(`{"Numero":1,"Saldo":200}`)
	jsonPuro := []byte(`{"n":1,"s":300}`)

	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)

}
