package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

type Conn interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

var DBConnection *sql.DB

func NewDB() error {
	sqlite3, err := sql.Open("sqlite3", "file:db.sqlite3?_foreign_keys=off")
	if err != nil {
		log.Printf("%q\n", err)
		return err
	}

	if err != nil {
		log.Printf("%q\n", err)
		return err
	}

	err = migrateDB(sqlite3)
	if err != nil {
		log.Printf("%q\n", err)
		return err
	}

	DBConnection = sqlite3

	return nil
}

func GetDBConnection() *sql.DB {
	return DBConnection
}

func migrateDB(db *sql.DB) error {
	sqlStmts := []string{
		`
	DROP TABLE IF EXISTS "orders";
	CREATE TABLE "orders" (
		id VARCHAR(64) NOT NULL PRIMARY KEY,
		total_price INT NOT NULL,
		status VARCHAR(64) CHECK( status IN ('WAITING', 'COMPLETE', 'CANCELED', 'RETURNED') ) NOT NULL
	);
	`,
		`
	DROP TABLE IF EXISTS "products";
	CREATE TABLE "products" (
		id VARCHAR(64) NOT NULL PRIMARY KEY,
		name VARCHAR(64) NOT NULL,
		description TEXT NOT NULL,
		price INT NOT NULL,
		category VARCHAR(64) NOT NULL,
		stock INT NOT NULL
	);
	`,
		`
	DROP TABLE IF EXISTS "order_products";
	CREATE TABLE "order_products" (
		order_id VARCHAR(64) NOT NULL,
		product_id VARCHAR(64) NOT NULL,
		amount INT NOT NULL,
		PRIMARY KEY (order_id, product_id),
		FOREIGN KEY (order_id) REFERENCES orders(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	);
	`,
	}

	for _, sqlStmt := range sqlStmts {
		if _, err := db.Exec(sqlStmt); err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return err
		}
	}

	return nil
}
