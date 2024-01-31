package product

import (
	"context"
	"database/sql"
	productDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"github.com/poster-keisuke/sample-clearn-architecture/app/usacase/transaction"
	"golang.org/x/xerrors"
)

type UpdateProductUseCase struct {
	productRepo productDomain.ProductRepository
	transaction transaction.Transaction
}

func NewUpdateProductUseCase(
	productRepo productDomain.ProductRepository,
	transaction transaction.Transaction,
) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepo: productRepo,
		transaction: transaction,
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

	if err := uc.transaction.StartTransaction(func(tx *sql.Tx) error {
		product.Name = input.Name
		product.Description = input.Description
		product.Price = input.Price
		product.Category = input.Category
		product.Stock = input.Stock

		err = uc.productRepo.Update(ctx, product)
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}
		return nil
	}); err != nil {
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
