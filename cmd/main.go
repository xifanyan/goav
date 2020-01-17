package main

import (
	"github.com/xifanyan/goav"
)

func main() {
	cfg := goav.DbConfig{
		Dialect: "sqlite3",
		// ConnectionStr: ":memory:",
		ConnectionStr: "./testdata/test.db",
	}

	database, _ := goav.NewDatabase(cfg)

	database.Init(&goav.Quote{})

	ts := goav.TimeSeries{
		Req:      goav.NewRequest(goav.TimeSeriesDailyAdjusted, "GDX", goav.IntervalNA, goav.FULL, goav.CSV, "60LA8W7JK9QYJEJN"),
		Database: database,
	}

	_ = ts.Save()

}
