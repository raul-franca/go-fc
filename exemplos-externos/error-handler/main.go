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
	err := errors.New("esta função simplesmente retorna um erro")
	return err
}
