package order

import (
	"context"
)

type OrderRepository interface {
	Create(ctx context.Context, order *Order) error
	GetByID(ctx context.Context, id string) (*Order, error)
	GetByIDAndStatus(ctx context.Context, orderID string, status OrderStatus) (*Order, error)
	Update(ctx context.Context, order *Order) error
}
