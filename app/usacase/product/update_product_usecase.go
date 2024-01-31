package product

import (
	"context"
	productDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"golang.org/x/xerrors"
)

type UpdateProductUseCase struct {
	productRepo productDomain.ProductRepository
}

func NewUpdateProductUseCase(
	productRepo productDomain.ProductRepository,
) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepo: productRepo,
	}
}

type UpdateProductUseCaseInputDto struct {
	ID          string
	Name        string
	Description string
	Price       int
	Category    string
	Stock       int
}

type UpdateProductUseCaseOutputDto struct {
	ID          string
	Name        string
	Description string
	Price       int
	Category    string
	Stock       int
}

func (uc *UpdateProductUseCase) Run(ctx context.Context, input UpdateProductUseCaseInputDto) (*UpdateProductUseCaseOutputDto, error) {
	product, err := uc.getValidProduct(ctx, input)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Category = input.Category
	product.Stock = input.Stock

	err = uc.productRepo.Update(ctx, product)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &UpdateProductUseCaseOutputDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Stock:       product.Stock,
	}, nil
}

func (uc *UpdateProductUseCase) getValidProduct(ctx context.Context, input UpdateProductUseCaseInputDto) (*productDomain.Product, error) {
	product, err := uc.productRepo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return product, nil
}
