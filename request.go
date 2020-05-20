package goav

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

type requestor interface {
	read() (io.ReadCloser, error)
}

// StockRequest ...
type StockRequest struct {
	gorm.Model
	Function        string     `gorm:"unique_index:idx_func_symb_intv"`
	Symbol          string     `gorm:"unique_index:idx_func_symb_intv"`
	Interval        string     `gorm:"unique_index:idx_func_symb_intv"`
	OutputSize      OutputSize `gorm:"-"`
	DataType        DataType   `gorm:"-"`
	APIKey          string     `gorm:"-"`
	StockTimeSeries []StockTimeSeries
}

// NewStockRequest ...
func NewStockRequest(function TimeSeriesFunction, symbol string, interval Interval, outputSize OutputSize,
	dataType DataType, apiKey string) *StockRequest {
	return &StockRequest{
		Function:   fmt.Sprint(function),
		Symbol:     symbol,
		Interval:   fmt.Sprint(interval),
		OutputSize: outputSize,
		DataType:   dataType,
		APIKey:     apiKey,
	}
}

func (req *StockRequest) getURL() string {
	var q []string

	q = append(q, fmt.Sprintf("function=%s", req.Function))
	q = append(q, fmt.Sprintf("apikey=%s", req.APIKey))
	q = append(q, fmt.Sprintf("symbol=%s", req.Symbol))
	q = append(q, fmt.Sprintf("datatype=%s", req.DataType))

	switch req.Function {
	case "TIME_SERIES_INTRADAY":
		q = append(q, fmt.Sprintf("interval=%s", req.Interval))
	case "TIME_SERIES_DAILY", "TIME_SERIES_DAILY_ADJUSTED":
		q = append(q, fmt.Sprintf("outputsize=%s", req.OutputSize))
	}

	return fmt.Sprintf("%s/query?%s", ALPHAVANTAGE, strings.Join(q, "&"))
}

func (req *StockRequest) read() (io.ReadCloser, error) {
	var err error

	url := req.getURL()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}
