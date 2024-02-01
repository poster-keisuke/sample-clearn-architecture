package order

import (
	"github.com/google/uuid"
	"github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"golang.org/x/xerrors"
)

type OrderStatus string

const (
	OrderStatusWaiting  OrderStatus = "WAITING"
	OrderStatusComplete OrderStatus = "COMPLETE"
	OrderStatusCanceled OrderStatus = "CANCELED"
	OrderStatusReturned OrderStatus = "RETURNED"
)

func (orderStatus OrderStatus) NeedsToRevertProductStock() bool {
	return orderStatus == OrderStatusCanceled || orderStatus == OrderStatusReturned
}

type OrderProcessType string

const (
	OrderProcessTypeShippingCompleted OrderProcessType = "SHIPPING_COMPLETED"
	OrderProcessTypeCanceled          OrderProcessType = "CANCELED"
	OrderProcessTypeReturned          OrderProcessType = "RETURNED"
)

func (orderProcessType OrderProcessType) Valid() error {
	switch orderProcessType {
	case OrderProcessTypeShippingCompleted, OrderProcessTypeCanceled, OrderProcessTypeReturned:
		return nil
	default:
		return xerrors.Errorf("invalid type: %s", orderProcessType)
	}
}

func (orderProcessType OrderProcessType) ValidTargetStatus(orderStatus OrderStatus) error {
	switch orderProcessType {
	case OrderProcessTypeShippingCompleted, OrderProcessTypeCanceled:
		if orderStatus != OrderStatusWaiting {
			return xerrors.Errorf("the order status must be WAITING: %s", orderStatus)
		}
	case OrderProcessTypeReturned:
		if orderStatus != OrderStatusComplete {
			return xerrors.Errorf("the order status must be COMPLETE: %s", orderStatus)
		}
	default:
		return xerrors.Errorf("unexpected process type: %s", orderProcessType)
	}
	return nil
}

func (orderProcessType OrderProcessType) UpdatedStatus() (OrderStatus, error) {
	switch orderProcessType {
	case OrderProcessTypeShippingCompleted:
		return OrderStatusComplete, nil
	case OrderProcessTypeCanceled:
		return OrderStatusCanceled, nil
	case OrderProcessTypeReturned:
		return OrderStatusReturned, nil
	default:
		return "", xerrors.Errorf("unexpected type: %s", orderProcessType)
	}
}

type Order struct {
	ID         string                    `json:"id"`
	Status     OrderStatus               `json:"status"`
	TotalPrice int                       `json:"total_price,omitempty"`
	Products   []*product.OrderedProduct `json:"products"`
}

func (o *Order) UpdateStatus(status OrderStatus) {
	o.Status = status
}

func (o *Order) UpdateProducts(products []*product.OrderedProduct) {
	o.Products = products
}

func (o *Order) UpdateTotalPrice(totalPrice int) {
	o.TotalPrice = totalPrice
}

func NewOrder() *Order {
	return &Order{
		ID:     uuid.NewString(),
		Status: OrderStatusWaiting,
	}
}
