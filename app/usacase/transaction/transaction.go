package transaction

import (
	"database/sql"
)

type Transaction interface {
	StartTransaction(fn func(ctx *sql.Tx) error) error
}
