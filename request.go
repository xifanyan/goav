package goav

import (
	"fmt"
	"strings"
)

type requestor interface {
	getURL() string
}

type StockRequest struct {
	Id         int        `gorm:"primary_key"`
	Function   Function   `gorm:"type:integer;unique_index:idx_func_symb_intv"`
	Symbol     string     `gorm:"unique_index:idx_func_symb_intv"`
	Interval   Interval   `gorm:"type:integer;unique_index:idx_func_symb_intv"`
	OutputSize OutputSize `gorm:"-"`
	DataType   DataType   `gorm:"-"`
	ApiKey     string     `gorm:"-"`
}

func NewStockRequest(function Function, symbol string, interval Interval, outputSize OutputSize,
	dataType DataType, apiKey string) *StockRequest {
	return &StockRequest{
		Function:   function,
		Symbol:     symbol,
		Interval:   interval,
		OutputSize: outputSize,
		DataType:   dataType,
		ApiKey:     apiKey,
	}
}

func (req *StockRequest) getURL() string {
	var q []string

	q = append(q, fmt.Sprintf("function=%s", req.Function))
	q = append(q, fmt.Sprintf("apikey=%s", req.ApiKey))
	q = append(q, fmt.Sprintf("symbol=%s", req.Symbol))
	q = append(q, fmt.Sprintf("datatype=%s", req.DataType))

	switch req.Function {
	case TimeSeriesIntraday:
		q = append(q, fmt.Sprintf("interval=%s", req.Interval))
	case TimeSeriesDaily, TimeSeriesDailyAdjusted:
		q = append(q, fmt.Sprintf("outputsize=%s", req.OutputSize))
	case TimeSeriesWeekly, TimeSeriesWeeklyAdjusted, TimeSeriesMonthly, TimeSeriesMonthlyAdjusted:
	}

	return fmt.Sprintf("%s/query?%s", ALPHAVANTAGE, strings.Join(q, "&"))
}
