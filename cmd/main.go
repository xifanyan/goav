package main

import (
	"log"

	"github.com/xifanyan/goav"
)

func main() {

	cfg := goav.DbConfig{
		Dialect:       "sqlite3",
		ConnectionStr: "./testdata/test.db",
	}

	db, _ := goav.NewDatabase(cfg)
	db.Init()

	symbols := []string{"GDX", "GLD", "SILJ", "SLV"}

	for _, symbol := range symbols {
		log.Printf("Processing %s ...", symbol)
		req := goav.NewStockRequest(goav.TimeSeriesDailyAdjusted, symbol, goav.NA, goav.FULL, goav.CSV, "61LA8W7JK9QYJEJN")
		ts := goav.NewStockTimeSeries(req, db)
		ts.Save()
	}

}
