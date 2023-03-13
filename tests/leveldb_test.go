package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

// Record represents a single record in the database
type Record struct {
	Timestamp int64
	Key       string
	Value     string
}

func benchmarkInsert(b *testing.B, db *leveldb.DB) {
	for i := 0; i < b.N; i++ {
		record := Record{
			Timestamp: time.Now().UnixNano(),
			Key:       strconv.Itoa(i),
			Value:     "value" + strconv.Itoa(i),
		}
		err := db.Put([]byte(strconv.FormatInt(record.Timestamp, 10)), []byte(record.Key+":"+record.Value), nil)
		if err != nil {
			b.Fatalf("Error inserting record: %v", err)
		}
	}
}

func TestInsert(t *testing.T) {
	db, err := leveldb.OpenFile("example.db", nil)
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	fmt.Println("\nInsert benchmark:")
	b := testing.Benchmark(func(b *testing.B) {
		benchmarkInsert(b, db)
	})
	t.Logf("\nInsert benchmark - %v", b)
}
