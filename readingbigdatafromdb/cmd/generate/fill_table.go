package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"math/rand"
	"time"
)

func fillTable(conn *pgx.Conn, num int) error {
	for i := 0; i < num; i++ {
		_, err := conn.Exec(context.Background(), `INSERT INTO "ledger_entries" (
    "amount1", "amount2", "amount3", "amount4", "amount5", "amount6", "amount7", "amount8", "amount9", "amount10",
	"date1", "date2", "date3", "date4", "date5", "date6", "date7", "date8", "date9", "date10") 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20);`,
			rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(),
			randDate(), randDate(), randDate(), randDate(), randDate(), randDate(), randDate(), randDate(), randDate(), randDate())
		if err != nil {
			return err
		}
		if i%1000 == 0 {
			fmt.Printf("Inserted %d rows\n", i)
		}
	}
	return nil
}

var minDate = time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
var maxDate = time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
var deltaDate = maxDate - minDate

func randDate() time.Time {
	sec := rand.Int63n(deltaDate) + minDate
	return time.Unix(sec, 0)
}
