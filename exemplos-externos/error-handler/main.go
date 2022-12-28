package main

import (
	"errors"
)

func main() {
	err := doesReturnError()
	if err != nil {
		panic(err)
	}
}
func doesReturnError() error {
	err := errors.New("Essa função simplesmente retorna um erro.")
	return err
}
