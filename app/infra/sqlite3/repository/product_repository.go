package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
	"github.com/poster-keisuke/sample-clearn-architecture/app/infra/sqlite3/db"
	"golang.org/x/xerrors"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository() product.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) GetByID(ctx context.Context, productID string) (*product.Product, error) {
	conn := db.GetDBConnection()
	stmt, err := conn.PrepareContext(ctx, "SELECT * FROM products WHERE id = ?;")
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	defer func() {
		_ = stmt.Close()
	}()

	row := stmt.QueryRowContext(ctx, productID)
	var p product.Product
	if err = row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Stock); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &p, nil
}

func (r *productRepository) Create(ctx context.Context, product *product.Product) error {
	conn := db.GetDBConnection()
	stmt, err := conn.PrepareContext(ctx, "INSERT INTO products (id, name, description, price, category, stock) VALUES (?, ?, ?, ?, ?, ?);")
	if err != nil {
		return fmt.Errorf(": %w", err)
	}
	defer func() {
		_ = stmt.Close()
	}()

	if _, err = stmt.ExecContext(ctx, product.ID, product.Name, product.Description, product.Price, product.Category, product.Stock); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func (r *productRepository) Update(ctx context.Context, product *product.Product) error {
	return nil
}
