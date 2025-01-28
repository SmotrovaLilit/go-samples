package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

const databaseURL = "postgres://postgres:postgres@localhost:2345/postgres?sslmode=disable"

func Connect(ctx context.Context) (*pgx.Conn, error, func() error) {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err, func() error { return nil }
	}
	return conn, nil, func() error {
		return conn.Close(ctx)
	}
}
