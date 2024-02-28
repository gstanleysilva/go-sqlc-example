package sqlc

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/gstanleysilva/go-sqlc-example/infra/database/gen"
)

type SQLCTxHelper struct {
	sqlConn *sql.DB
}

func NewSQLCHelper(dbConn *sql.DB) *SQLCTxHelper {
	return &SQLCTxHelper{
		sqlConn: dbConn,
	}
}

func (c *SQLCTxHelper) CallTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.sqlConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", rbErr, err)
		}
		return err
	}

	return tx.Commit()
}
