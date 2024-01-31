package db

import (
	"context"
	"database/sql"
	"golang.org/x/xerrors"
)

type txKey struct{}

func SetTxToCtx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func RemoveTxFromCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, txKey{}, nil)
}

func TxFromCtx(ctx context.Context) (*sql.Tx, error) {
	txAny := ctx.Value(txKey{})
	if txAny == nil {
		return nil, nil
	}
	tx, ok := txAny.(*sql.Tx)
	if !ok {
		return nil, xerrors.Errorf("unexpected Tx type: %T", txAny)
	}
	return tx, nil
}
