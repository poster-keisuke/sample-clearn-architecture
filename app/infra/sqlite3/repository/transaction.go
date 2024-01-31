package repository

import (
	"context"
	"database/sql"
	"github.com/poster-keisuke/sample-clearn-architecture/app/infra/sqlite3/db"
	"github.com/poster-keisuke/sample-clearn-architecture/app/usacase/transaction"
	"log"
)

type Transaction struct {
}

func NewTransaction() transaction.Transaction {
	return &Transaction{}
}

func (t *Transaction) StartTransaction(fn func(tx *sql.Tx) error) error {
	conn := db.GetDBConnection()
	ctx := context.Background()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	db.SetTxToCtx(ctx, tx)
	defer db.RemoveTxFromCtx(ctx)

	if err = fn(tx); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			log.Printf("%q\n", err)
		}
		return err
	}
	return tx.Commit()
}
