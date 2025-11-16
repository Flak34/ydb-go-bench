package main

import (
	"strings"
	"testing"

	"github.com/ydb-platform/ydb-go-sdk/v3/query"
)

func BenchmarkConcurrentResultSetsEnabled(b *testing.B) {
	q := strings.Join(selects, "")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := db.Query().Query(b.Context(), q, query.WithConcurrentResultSets(true))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkConcurrentResultSetsDisabled(b *testing.B) {
	q := strings.Join(selects, "")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := db.Query().Query(b.Context(), q, query.WithConcurrentResultSets(false))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMultipleRequests(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range selects {
			_, err := db.Query().Query(b.Context(), q)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}

// Первый прогон
//BenchmarkConcurrentResultSetsEnabled
//BenchmarkConcurrentResultSetsEnabled-12     	      34	  36461410 ns/op
//BenchmarkConcurrentResultSetsDisabled
//BenchmarkConcurrentResultSetsDisabled-12    	      31	  37751904 ns/op
//BenchmarkMultipleRequests
//BenchmarkMultipleRequests-12                	       1	1017505757 ns/op

// Второй прогон
//BenchmarkConcurrentResultSetsEnabled
//BenchmarkConcurrentResultSetsEnabled-12     	      19	  64236576 ns/op
//BenchmarkConcurrentResultSetsDisabled
//BenchmarkConcurrentResultSetsDisabled-12    	       3	 419940854 ns/op
//BenchmarkMultipleRequests
//BenchmarkMultipleRequests-12                	       1	1578824543 ns/op
