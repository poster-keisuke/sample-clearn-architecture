package order

import (
	"context"
	"database/sql"
	orderDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/order"
	productDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"github.com/poster-keisuke/sample-clearn-architecture/app/usacase/transaction"
	"golang.org/x/xerrors"
)

type CancelOrderUseCase struct {
	orderRepo   orderDomain.OrderRepository
	productRepo productDomain.ProductRepository
	transaction transaction.Transaction
}

func (uc *CancelOrderUseCase) Run(ctx context.Context, orderID string) error {
	order, err := uc.orderRepo.GetByIDAndStatus(ctx, orderID, orderDomain.OrderStatusWaiting)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	if err := uc.transaction.StartTransaction(func(tx *sql.Tx) error {

		order.UpdateStatus(orderDomain.OrderStatusCanceled)
		if err := uc.orderRepo.Update(ctx, order); err != nil {
			return xerrors.Errorf(": %w", err)
		}

		products, err := getRevertProductStock(ctx, order)
		if err != nil {
			xerrors.Errorf(": %w", err)
		}

		for _, product := range products {
			if err := uc.productRepo.Update(ctx, product); err != nil {
				return xerrors.Errorf(": %w", err)
			}
		}

		return nil
	}); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
