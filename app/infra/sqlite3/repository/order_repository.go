package repository

import (
	"context"
	"database/sql"
	"github.com/poster-keisuke/sample-clearn-architecture/app/domain/order"
	"github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"github.com/poster-keisuke/sample-clearn-architecture/app/infra/sqlite3/db"
	"golang.org/x/xerrors"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository() order.OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) conn(ctx context.Context) (db.Conn, error) {
	tx, err := db.TxFromCtx(ctx)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	if tx == nil {
		return r.db, nil
	}
	return tx, nil
}

func (r *orderRepository) GetByID(ctx context.Context, orderID string) (*order.Order, error) {
	return nil, nil
}

func (r *orderRepository) GetByIDAndStatus(ctx context.Context, orderID string, status order.OrderStatus) (*order.Order, error) {
	conn, err := r.conn(ctx)
	stmt, err := conn.PrepareContext(ctx, `
SELECT o.id,
       o.total_price,
       o.status,
       p.id,
       p.name,
       p.description,
       p.price,
       p.category,
       op.amount 
FROM orders o
INNER JOIN order_products op ON o.id = op.order_id
INNER JOIN products p ON p.id = op.product_id
WHERE o.id = ? AND o.status = ?;`)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	defer func() {
		_ = stmt.Close()
	}()

	rows, err := stmt.QueryContext(ctx, orderID, status)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	var o order.Order
	for rows.Next() {
		var (
			p      product.Product
			amount int
		)
		if err = rows.Scan(
			&o.ID,
			&o.TotalPrice,
			&o.Status,
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Category,
			&amount,
		); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		o.Products = append(o.Products, product.NewProductWithOrderedAmount(&p, amount))
	}

	return &o, nil
}

func (r *orderRepository) Create(ctx context.Context, order *order.Order) error {
	return nil
}

func (r *orderRepository) Update(ctx context.Context, order *order.Order) error {
	conn, err := r.conn(ctx)
	stmt, err := conn.PrepareContext(ctx, "UPDATE orders SET total_price = ?, status = ? WHERE id = ?;")
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	defer func() {
		_ = stmt.Close()
	}()

	if _, err = stmt.ExecContext(ctx, order.TotalPrice, order.Status, order.ID); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}
