package main

import (
	"github.com/xifanyan/goav"
)

func main() {
	cfg := goav.DbConfig{
		Dialect:       "sqlite3",
		ConnectionStr: "./testdata/test.db",
	}

	req := goav.NewStockRequest(goav.TimeSeriesDailyAdjusted, "GOOG", goav.IntervalNA, goav.FULL, goav.CSV, "demo")

	database, _ := goav.NewDatabase(cfg)
	database.Init()

	_ = goav.NewTimeSeries(req, database)
	//
}
