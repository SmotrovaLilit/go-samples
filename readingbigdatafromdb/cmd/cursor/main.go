package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"readingbigdatafromdb/internal/db"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %v\n", time.Since(now))
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err, cleanup := db.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		panic(err)
	}

	query := `SELECT * FROM "ledger_entries"`

	// Declare cursor for query.
	_, err = tx.Exec(context.Background(), `DECLARE cur CURSOR FOR `+query)
	if err != nil {
		panic(err)
	}

	limit := 1000
	readCount := 0

	// read all rows in 1000 row chunks
	start := time.Now()
	defer func() {
		fmt.Println("Time taken: ", time.Since(start))
	}()
	for i := 0; ; i++ {
		rows, err := tx.Query(context.Background(), fmt.Sprintf(`FETCH %d FROM cur`, limit))
		if err != nil {
			panic(err)
		}
		cnt := 0
		for rows.Next() {
			cnt++
			_, err = rows.Values()
			if err != nil {
				panic(err)
			}
		}
		if cnt == 0 {
			break
		}
		readCount += cnt
		if readCount%100_000 == 0 {
			fmt.Printf("Read %d rows\n", readCount)
		}
		if readCount == 1_000_000 {
			break
		}
	}

	if _, err = tx.Exec(context.Background(), `CLOSE cur`); err != nil {
		log.Fatalf("Ошибка закрытия курсора: %v", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		panic(err)
	}
}
