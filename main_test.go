package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ydb-platform/ydb-go-sdk/v3"
)

const (
	selectQuery = "SELECT * FROM test1 WHERE test1.id > %d limit 10;"
	rowsCount   = 1000000000
	selectCount = 20
)

var (
	db      *ydb.Driver
	selects []string
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	db = connectDB()

	selects = makeSelects(selectCount)

	code := m.Run()

	db.Close(ctx)

	os.Exit(code)
}

func makeSelects(count int) []string {
	var res []string
	for i := range count {
		res = append(res, fmt.Sprintf(selectQuery, rowsCount/count*i))
	}

	return res
}
