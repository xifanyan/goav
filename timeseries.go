package goav

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/xifanyan/csvstream"
)

type StockTimeSeries struct {
	gorm.Model
	req              *StockRequest `gorm:"-"`
	db               *Database     `gorm:"-"`
	StockRequestID   uint          `gorm:"unique_index:idx_id_timestamp"`
	TimeStamp        string        `gorm:"unique_index:idx_id_timestamp" csv:"timestamp"`
	Open             float64       `csv:"open"`
	High             float64       `csv:"high"`
	Low              float64       `csv:"low"`
	Close            float64       `csv:"close"`
	AdjustedClose    float64       `csv:"adjusted_close"`
	Volume           float64       `csv:"volume"`
	Dividend         float64       `csv:"dividend_amount"`
	SplitCoefficient float64       `csv:"split_coefficient"`
}

func NewStockTimeSeries(req *StockRequest, db *Database) *StockTimeSeries {
	return &StockTimeSeries{
		req: req,
		db:  db,
	}
}

func (sts *StockTimeSeries) Save() error {
	var err error

	err = sts.db.inst.Where(&StockRequest{Function: sts.req.Function, Symbol: sts.req.Symbol, Interval: sts.req.Interval}).FirstOrCreate(sts.req).Error
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	rc, err := sts.req.read()
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	defer rc.Close()

	dec, err := csvstream.NewDecoder(rc, &StockTimeSeries{})
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	c, err := dec.Unmarshal()
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	for row := range c {
		ts := row.(StockTimeSeries)
		sts.db.inst.Where(&StockTimeSeries{TimeStamp: ts.TimeStamp, StockRequestID: sts.req.ID}).Assign(StockTimeSeries{Open: ts.Open, High: ts.High, Low: ts.Low, Close: ts.Close, AdjustedClose: ts.AdjustedClose, Volume: ts.Volume, Dividend: ts.Dividend, SplitCoefficient: ts.SplitCoefficient}).FirstOrCreate(&ts)
		// sts.db.inst.Model(sts.req).Association("StockTimeSeries").Append(ts)
	}

	return err
}
