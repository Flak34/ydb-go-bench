package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"path"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
	yc "github.com/ydb-platform/ydb-go-yc"
)

func connectDB() *ydb.Driver {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := ydb.Open(ctx,
		"grpcs://lb.etnmotudqemrlhj1r5ec.ydb.mdb.yandexcloud.net:2135/ru-central1/b1gijrv5hn39vir6rfva/etnmotudqemrlhj1r5ec",
		yc.WithInternalCA(),
		yc.WithServiceAccountKeyFileCredentials("./key.json"),
	)
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	ctx := context.TODO()
	db := connectDB()
	defer db.Close(ctx)

	fillTable(db, "test1")
}

func makeTable(db *ydb.Driver, tableName string) {
	err := db.Table().Do(context.Background(),
		func(ctx context.Context, s table.Session) (err error) {
			return s.CreateTable(ctx, path.Join(db.Name(), tableName),
				options.WithColumn("series_id", types.TypeUint64), // not null column
				options.WithColumn("title", types.Optional(types.TypeUTF8)),
				options.WithColumn("series_info", types.Optional(types.TypeUTF8)),
				options.WithColumn("release_date", types.Optional(types.TypeDate)),
				options.WithColumn("comment", types.Optional(types.TypeUTF8)),
				options.WithPrimaryKeyColumn("series_id"),
			)
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fillTable(db, tableName)
}

func fillTable(db *ydb.Driver, tableName string) {
	const total = 100000000
	const batchSize = 50000

	rand.Seed(time.Now().UnixNano())

	titles := []string{"Star Wars", "Breaking Bad", "Game of Thrones", "Friends", "The Office", "Stranger Things"}
	comments := []string{"great", "boring", "funny", "awesome", "classic", "underrated"}
	infoSamples := []string{"Sci-Fi", "Drama", "Comedy", "Fantasy", "Thriller", "Action"}

	start := time.Now()
	log.Printf("Начинаем вставку %d строк...", total)

	for i := 0; i < total; i += batchSize {
		var (
			rows []types.Value
		)

		for j := 0; j < batchSize && i+j < total; j++ {
			id := uint64(i + j + 1)
			title := titles[rand.Intn(len(titles))]
			info := infoSamples[rand.Intn(len(infoSamples))]
			comment := comments[rand.Intn(len(comments))]
			releaseDate := time.Date(
				rand.Intn(20)+2000,
				time.Month(rand.Intn(12)+1),
				rand.Intn(28)+1,
				0, 0, 0, 0,
				time.UTC,
			)

			rows = append(rows, types.StructValue(
				types.StructFieldValue("id", types.Int64Value(int64(id))),
				types.StructFieldValue("title", types.UTF8Value(title)),
				types.StructFieldValue("series_info", types.UTF8Value(info)),
				types.StructFieldValue("release_date", types.DateValueFromTime(releaseDate)),
				types.StructFieldValue("comment", types.UTF8Value(comment)),
			))
		}

		err := db.Table().BulkUpsert(context.Background(),
			"/ru-central1/b1gijrv5hn39vir6rfva/etnmotudqemrlhj1r5ec/test1",
			table.BulkUpsertDataRows(types.ListValue(rows...)),
		)
		if err != nil {
			log.Fatalf("Ошибка вставки: %v", err)
		}

		if (i/batchSize)%10 == 0 {
			log.Printf("Вставлено %d строк...", i)
		}
	}

	log.Printf("Готово! Вставлено %d строк за %v", total, time.Since(start))

}
