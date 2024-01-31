package product

import (
	"context"
)

type ProductRepository interface {
	GetByID(ctx context.Context, productID string) (*Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
}
