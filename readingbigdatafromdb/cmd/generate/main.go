package main

import (
	"context"
	"readingbigdatafromdb/internal/db"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err, cleanup := db.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	err = createTable(conn)
	if err != nil {
		panic(err)
	}

	err = fillTable(conn, 8_000_000)
	if err != nil {
		panic(err)
	}
}
