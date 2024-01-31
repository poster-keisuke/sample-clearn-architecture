package repository

import (
	"context"
	"github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
)

type productRepository struct {
}

func NewProductRepository() product.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) GetByID(ctx context.Context, productID string) (*product.Product, error) {
	return nil, nil
}

func (r *productRepository) Create(ctx context.Context, product *product.Product) error {
	return nil
}

func (r *productRepository) Update(ctx context.Context, product *product.Product) error {
	return nil
}
