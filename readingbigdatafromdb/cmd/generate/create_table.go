package main

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func createTable(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS "ledger_entries" (
    "id" SERIAL PRIMARY KEY,
    "amount1" NUMERIC(98,20) NOT NULL,
    "amount2" NUMERIC(98,20) NOT NULL,
    "amount3" NUMERIC(98,20) NOT NULL,
    "amount4" NUMERIC(98,20) NOT NULL,
    "amount5" NUMERIC(98,20) NOT NULL,
    "amount6" NUMERIC(98,20) NOT NULL,
    "amount7" NUMERIC(98,20) NOT NULL,
    "amount8" NUMERIC(98,20) NOT NULL,
    "amount9" NUMERIC(98,20) NOT NULL,
    "amount10" NUMERIC(98,20) NOT NULL,
    "date1" TIMESTAMP NOT NULL,
    "date2" TIMESTAMP NOT NULL,
    "date3" TIMESTAMP NOT NULL,
    "date4" TIMESTAMP NOT NULL,
    "date5" TIMESTAMP NOT NULL,
    "date6" TIMESTAMP NOT NULL,
    "date7" TIMESTAMP NOT NULL,
    "date8" TIMESTAMP NOT NULL,
    "date9" TIMESTAMP NOT NULL,
    "date10" TIMESTAMP NOT NULL
);`)
	return err
}
