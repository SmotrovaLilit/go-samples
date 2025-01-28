package main

import (
	"context"
	"fmt"
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

	start := time.Now()
	defer func() {
		fmt.Println("Time taken: ", time.Since(start))
	}()
	rows, err := conn.Query(context.Background(), `SELECT * FROM "ledger_entries"`)
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
		if cnt%100_000 == 0 {
			fmt.Printf("Read %d rows\n", cnt)
		}
		if cnt == 1_000_000 {
			break
		}
	}
}
