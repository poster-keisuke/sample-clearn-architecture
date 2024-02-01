package order

import (
	"context"
	"database/sql"
	orderDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/order"
	productDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"github.com/poster-keisuke/sample-clearn-architecture/app/usacase/transaction"
	"golang.org/x/xerrors"
)

type CreateOrderUseCase struct {
	orderRepo   orderDomain.OrderRepository
	productRepo productDomain.ProductRepository
	transaction transaction.Transaction
}

func NewCreteOrderUseCase(
	orderRepo orderDomain.OrderRepository,
	productRepo productDomain.ProductRepository,
	transaction transaction.Transaction,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
		transaction: transaction,
	}
}

type ProductIDAndAmount struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

type CreteOrderUseCaseInputDto struct {
	Products []*ProductIDAndAmount `json:"products"`
}

func (uc *CreateOrderUseCase) Run(ctx context.Context, input CreteOrderUseCaseInputDto) error {
	o := orderDomain.NewOrder()

	productIDs := make([]string, 0, len(input.Products))
	for _, p := range input.Products {
		productIDs = append(productIDs, p.ID)
	}

	products, err := uc.productRepo.GetMultiByIDs(ctx, productIDs)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	productMap := make(map[string]*productDomain.Product)
	for _, p := range products {
		productMap[p.ID] = p
	}

	var orderedProducts []*productDomain.OrderedProduct
	for _, p := range input.Products {
		product, ok := productMap[p.ID]
		if !ok {
			return xerrors.Errorf("product_id is not found: %s", p.ID)
		}
		orderedProducts = append(orderedProducts, productDomain.NewOrderedProduct(product, p.Amount))
		if err := product.DecreaseStockFromAmount(p.Amount); err != nil {
			return xerrors.Errorf(": %w", err)
		}

	}

	totalPrice := summaryTotalPrice(orderedProducts)
	o.UpdateProducts(orderedProducts)
	o.UpdateTotalPrice(totalPrice)

	if err := uc.transaction.StartTransaction(func(tx *sql.Tx) error {

		if err = uc.orderRepo.Create(ctx, o); err != nil {
			return xerrors.Errorf(": %w", err)
		}

		for _, product := range products {
			if err = uc.productRepo.Update(ctx, product); err != nil {
				return xerrors.Errorf(": %w", err)
			}
		}

		return nil

	}); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func summaryTotalPrice(orderedProducts []*productDomain.OrderedProduct) int {
	var totalPrice int
	for _, product := range orderedProducts {
		totalPrice += product.CalculatePrice(product.Price)
	}
	return totalPrice
}
