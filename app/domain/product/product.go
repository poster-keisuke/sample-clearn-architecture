package product

import (
	"fmt"
	"github.com/google/uuid"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Stock       int    `json:"stock"`
}
type Products []*Product

type OrderedProduct struct {
	*Product
	Amount int `json:"amount"`
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

func NewOrderedProduct(product *Product, amount int) *OrderedProduct {
	return &OrderedProduct{
		Product: product,
		Amount:  amount,
	}
}

func (p *Product) RevertStockFromAmount(amount int) {
	p.Stock += amount
}

func (p *Product) DecreaseStockFromAmount(amount int) error {
	p.Stock -= amount
	if p.Stock < 0 {
		return fmt.Errorf("the product's stock is not enough: id: %s, stock: %d", p.ID, p.Stock)
	}
	return nil
}

func (ps Products) MapByID() map[string]*Product {
	m := make(map[string]*Product, len(ps))
	for _, p := range ps {
		m[p.ID] = p
	}
	return m
}

func (ps Products) RevertStockFromProductsWithOrderedAmount(orderedProducts []*OrderedProduct) error {
	m := ps.MapByID()
	for _, orderedProduct := range orderedProducts {
		p, ok := m[orderedProduct.ID]
		if !ok {
			return fmt.Errorf("product_id is not found: %s", orderedProduct.ID)
		}
		p.RevertStockFromAmount(orderedProduct.Amount)
	}
	return nil
}

func (ps Products) DecreaseStockFromOrderedProduct(orderedProducts []*OrderedProduct) error {
	m := ps.MapByID()
	for _, orderedProduct := range orderedProducts {
		p, ok := m[orderedProduct.ID]
		if !ok {
			return fmt.Errorf("product_id is not found: %s", orderedProduct.ID)
		}
		if err := p.DecreaseStockFromAmount(orderedProduct.Amount); err != nil {
			return err
		}
	}
	return nil
}

func (orderedProduct *OrderedProduct) CalculatePrice(price int) int {
	return orderedProduct.Amount * price
}
