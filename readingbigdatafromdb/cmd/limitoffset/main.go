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

	limit := 1000

	query := `SELECT * FROM "ledger_entries" LIMIT %d OFFSET %d`
	readCount := 0
	// read all rows in 1000 row chunks
	start := time.Now()
	defer func() {
		fmt.Println("Time taken: ", time.Since(start))
	}()
	for i := 0; ; i++ {
		offset := i * limit
		rows, err := conn.Query(context.Background(), fmt.Sprintf(query, limit, offset))
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
}
