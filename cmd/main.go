package main

import (
	"github.com/golang/glog"
	"github.com/xifanyan/goav"
)

func main() {
	cfg := goav.DbConfig{
		Dialect: "sqlite3",
		// ConnectionStr: ":memory:",
		ConnectionStr: "./testdata/test.db",
	}

	db, _ := goav.NewDatabase(cfg)

	db.DB.AutoMigrate(&goav.Quote{})

	if !db.DB.HasTable(&goav.Quote{}) {
		glog.Error("No Table")
	}

	ts := goav.TimeSeries{
		Req:      goav.NewRequest(goav.TimeSeriesDailyAdjusted, "GDX", goav.IntervalNA, goav.FULL, goav.CSV, "60LA8W7JK9QYJEJN"),
		Database: db,
	}

	_ = ts.Save()

}
