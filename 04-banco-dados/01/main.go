package main

import "github.com/google/uuid"
//_ "github.com/go-sql-driver/mysql"
type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {

}
