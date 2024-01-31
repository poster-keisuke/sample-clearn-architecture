package product

import (
	"context"
	productDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"golang.org/x/xerrors"
)

type CreateProductUseCase struct {
	productRepo productDomain.ProductRepository
}

func NewCreateProductUseCase(
	productRepo productDomain.ProductRepository,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepo: productRepo,
	}
}

type CreateProductUseCaseInputDto struct {
	Name        string
	Description string
	Price       int
	Category    string
	Stock       int
}

type CreateProductUseCaseOutputDto struct {
	ID          string
	Name        string
	Description string
	Price       int
	Category    string
	Stock       int
}

func (uc *CreateProductUseCase) Run(ctx context.Context, input CreateProductUseCaseInputDto) (*CreateProductUseCaseOutputDto, error) {
	product := productDomain.NewProduct(input.Name, input.Description, input.Category, input.Price, input.Stock)

	err := uc.productRepo.Create(ctx, product)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &CreateProductUseCaseOutputDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Stock:       product.Stock,
	}, nil
}
