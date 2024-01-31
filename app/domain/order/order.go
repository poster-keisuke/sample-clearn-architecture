package order

import "github.com/google/uuid"

type OrderStatus string

const (
	OrderStatusWaiting  OrderStatus = "WAITING"
	OrderStatusComplete OrderStatus = "COMPLETE"
	OrderStatusCanceled OrderStatus = "CANCELED"
	OrderStatusReturned OrderStatus = "RETURNED"
)

type Order struct {
	ID         string           `json:"id"`
	Status     OrderStatus      `json:"status"`
	TotalPrice int              `json:"total_price,omitempty"`
	Products   []OrderedProduct `json:"products"`
}

type OrderedProduct struct {
	OrderID   string `json:"order_id"`
	ProductID string `json:"productID"`
	Amount    int    `json:"amount"`
}

func NewOrder() *Order {
	return &Order{
		ID:     uuid.NewString(),
		Status: OrderStatusWaiting,
	}
}
