package main

import (
	"github.com/xifanyan/goav"
)

func main() {
	cfg := goav.DbConfig{
		Dialect:       "sqlite3",
		ConnectionStr: "./testdata/test.db",
	}

	db, _ := goav.NewDatabase(cfg)
	db.Init()

	req1 := goav.NewStockRequest(goav.TimeSeriesDailyAdjusted, "GDX", goav.NA, goav.FULL, goav.CSV, "61LA8W7JK9QYJEJN")
	ts1 := goav.NewStockTimeSeries(req1, db)
	ts1.Save()

	req2 := goav.NewStockRequest(goav.TimeSeriesDailyAdjusted, "GLD", goav.NA, goav.FULL, goav.CSV, "61LA8W7JK9QYJEJN")
	ts2 := goav.NewStockTimeSeries(req2, db)
	ts2.Save()

	req3 := goav.NewStockRequest(goav.TimeSeriesDailyAdjusted, "SILJ", goav.NA, goav.FULL, goav.CSV, "61LA8W7JK9QYJEJN")
	ts3 := goav.NewStockTimeSeries(req3, db)
	ts3.Save()

	req4 := goav.NewStockRequest(goav.TimeSeriesDailyAdjusted, "SLV", goav.NA, goav.FULL, goav.CSV, "61LA8W7JK9QYJEJN")
	ts4 := goav.NewStockTimeSeries(req4, db)
	ts4.Save()
}
