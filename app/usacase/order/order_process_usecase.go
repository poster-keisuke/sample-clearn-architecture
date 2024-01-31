package order

import (
	"context"
	"database/sql"
	orderDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/order"
	productDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"github.com/poster-keisuke/sample-clearn-architecture/app/usacase/transaction"
	"golang.org/x/xerrors"
)

type OrderProcessUseCase struct {
	orderRepo   orderDomain.OrderRepository
	productRepo productDomain.ProductRepository
	transaction transaction.Transaction
}

func (uc *OrderProcessUseCase) Run(ctx context.Context, orderID string, orderProcessType orderDomain.OrderProcessType) (*orderDomain.Order, error) {
	var order *orderDomain.Order
	if err := orderProcessType.Valid(); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	order, err := uc.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	if err = orderProcessType.ValidTargetStatus(order.Status); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	updatedStatus, err := orderProcessType.UpdatedStatus()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	order.UpdateStatus(updatedStatus)

	if err := uc.transaction.StartTransaction(func(tx *sql.Tx) error {
		if err := uc.orderRepo.Update(ctx, order); err != nil {
			return xerrors.Errorf(": %w", err)
		}

		if updatedStatus.NeedsToRevertProductStock() {
			// TODO: revert product stock
		}

		return nil
	}); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return order, nil
}
