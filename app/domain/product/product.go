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

func Reconstruct(
	id,
	name,
	description,
	category string,
	price int,
	stock int,
) *Product {
	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Stock:       stock,
	}
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

type OrderedProduct struct {
	*Product
	Amount int `json:"amount"`
}

func NewOrderedProduct(product *Product, amount int) *OrderedProduct {
	return &OrderedProduct{
		Product: product,
		Amount:  amount,
	}
}
