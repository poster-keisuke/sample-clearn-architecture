package order

import (
	"context"
	"github.com/poster-keisuke/sample-clearn-architecture/app/domain/order"
	"github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"golang.org/x/xerrors"
)

func getRevertProductStock(ctx context.Context, order *order.Order) (product.Products, error) {
	var products product.Products
	for _, p := range order.Products {
		products = append(products, product.Reconstruct(p.ID, p.Name, p.Description, p.Category, p.Price, p.Stock))
	}

	if err := products.RevertStockFromProductsWithOrderedAmount(order.Products); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return products, nil
}
