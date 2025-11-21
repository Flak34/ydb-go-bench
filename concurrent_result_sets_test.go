package main

import (
	"strings"
	"testing"

	"github.com/ydb-platform/ydb-go-sdk/v3/query"
)

// Bench с включенной опцией
func BenchmarkConcurrentResultSetsEnabled(b *testing.B) {
	q := strings.Join(selects, "")
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = db.Query().Query(b.Context(), q, query.WithConcurrentResultSets(true))
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Bench с выключенной опцией
func BenchmarkConcurrentResultSetsDisabled(b *testing.B) {
	q := strings.Join(selects, "")
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = db.Query().Query(b.Context(), q, query.WithConcurrentResultSets(false))
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Просто для сравнения серия последовательных запросов
func BenchmarkMultipleRequests(b *testing.B) {
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, q := range selects {
			_, err = db.Query().Query(b.Context(), q)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
