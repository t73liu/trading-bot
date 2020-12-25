package traderdb

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Interface for pgx.Conn and pgxpool.Pool
type PGConnection interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}
