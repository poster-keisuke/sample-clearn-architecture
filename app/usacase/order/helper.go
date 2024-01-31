package order

import (
	"context"
	"database/sql"
	"github.com/poster-keisuke/sample-clearn-architecture/app/domain/order"
)

func revertProductStockTx(ctx context.Context, tx *sql.Tx, order *order.Order) error {
	//for _, product := range order.Products {
	//	product.Stock += product.OrderedAmount
	//	if err := product.UpdateTx(ctx, tx); err != nil {
	//		return err
	//	}
	//}
	//order.Products

	//products := domain.NewProductsFromProductsWithOrderedAmount(order.Products)
	//if err := products.RevertStockFromProductsWithOrderedAmount(order.Products); err != nil {
	//	return xerrors.Errorf(": %w", err)
	//}
	//
	//for _, product := range products {
	//	if err := u.db.Product.UpdateTx(ctx, tx, product); err != nil {
	//		return xerrors.Errorf(": %w", err)
	//	}
	//}

	return nil
}
