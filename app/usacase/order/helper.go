package order

import (
	"context"
	"github.com/poster-keisuke/sample-clearn-architecture/app/domain/order"
	"github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
)

func getRevertProductStock(ctx context.Context, order *order.Order) ([]*product.Product, error) {
	var products []*product.Product
	for _, p := range order.Products {
		products = append(products, product.Reconstruct(p.ID, p.Name, p.Description, p.Category, p.Price, p.Stock))
	}

	//products := domain.NewProductsFromProductsWithOrderedAmount(order.Products)
	//if err := products.RevertStockFromProductsWithOrderedAmount(order.Products); err != nil {
	//	return xerrors.Errorf(": %w", err)
	//}
	//

	return products, nil
}
