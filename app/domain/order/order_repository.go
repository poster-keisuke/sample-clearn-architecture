package order

import "context"

type OrderRepository interface {
	GetByID(ctx context.Context, id string) (*Order, error)
	Create(ctx context.Context, order *Order) error
	Update(ctx context.Context, order *Order) error
}
