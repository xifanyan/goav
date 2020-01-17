package main

import (
	"github.com/xifanyan/goav"
)

func main() {
	cfg := goav.DbConfig{
		Dialect:       "sqlite3",
		ConnectionStr: "./testdata/test.db",
	}

	database, _ := goav.NewDatabase(cfg)

	database.Init(&goav.Quote{})

	ts := goav.TimeSeries{
		Req:      goav.NewRequest(goav.TimeSeriesDailyAdjusted, "GOOG", goav.IntervalNA, goav.FULL, goav.CSV, "demo"),
		Database: database,
	}

	_ = ts.Save()

}
