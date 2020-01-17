package goav

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/xifanyan/csvstream"
)

// Quote ...
type Quote struct {
	TimeStamp        string  `gorm:"primary_key" csv:"timestamp"`
	Function         uint8   `gorm:"primary_key;auto_increment:false" csv:"function"`
	Symbol           string  `gorm:"primary_key"`
	Open             float64 `csv:"open"`
	High             float64 `csv:"high"`
	Low              float64 `csv:"low"`
	Close            float64 `csv:"close"`
	AdjustedClose    float64 `csv:"adjusted_close"`
	Volume           int     `csv:"volume"`
	Dividend         float64 `csv:"dividend_amount"`
	SplitCoefficient float64 `csv:"split_coefficient"`
}

// TimeSeries ...
type TimeSeries struct {
	Req      *Request
	Database *Database
}

// NewTimeSeries : constructor for timeseries
func NewTimeSeries(db *Database, req *Request) (*TimeSeries, error) {
	return &TimeSeries{
		Req:      req,
		Database: db,
	}, nil
}

// Save : save timeseires to database
func (ts *TimeSeries) Save() error {
	url := ts.Req.GetURL()

	resp, err := http.Get(url)
	if err != nil {
		glog.Errorf("Failed to fetch file from URL %s", url)
		return err
	}
	defer resp.Body.Close()

	dec, err := csvstream.NewDecoder(resp.Body, &Quote{})
	if err != nil {
		glog.Error("Failed to initialize decoder")
		return err
	}

	c, err := dec.Unmarshal()
	if err != nil {
		return err
	}

	for row := range c {
		quote := &Quote{}
		q := row.(Quote)
		q.Function = uint8(ts.Req.function)
		q.Symbol = ts.Req.symbol
		//ts.Database.DB.Create(q)
		//ts.Database.DB.Debug().FirstOrCreate(&quote, q)
		ts.Database.DB.Where(Quote{TimeStamp: q.TimeStamp, Function: q.Function, Symbol: q.Symbol}).Attrs(q).FirstOrCreate(&quote)
	}

	return err
}
