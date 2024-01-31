package product

import "github.com/google/uuid"

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Stock       int    `json:"stock"`
}

func NewProduct(name, description, category string, price, stock int) *Product {
	return &Product{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Stock:       stock,
	}
}

type WithOrderedAmount struct {
	*Product
	Amount int `json:"amount"`
}

func NewProductWithOrderedAmount(product *Product, amount int) *WithOrderedAmount {
	return &WithOrderedAmount{
		Product: product,
		Amount:  amount,
	}
}
