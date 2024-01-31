package product

import (
	"context"
	"fmt"
	productDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"golang.org/x/xerrors"
)

type GetProductUseCase struct {
	productRepo productDomain.ProductRepository
}

func NewGetProductUseCase(
	productRepo productDomain.ProductRepository,
) *GetProductUseCase {
	return &GetProductUseCase{
		productRepo: productRepo,
	}
}

type GetProductUseCaseOutputDto struct {
	ID          string
	Name        string
	Description string
	Price       int
	Category    string
	Stock       int
}

func (uc *GetProductUseCase) Run(ctx context.Context, id string) (*GetProductUseCaseOutputDto, error) {
	fmt.Println("GetProductUseCase")
	product, err := uc.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &GetProductUseCaseOutputDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Stock:       product.Stock,
	}, nil
}
