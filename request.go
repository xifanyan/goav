package goav

import (
	"fmt"
	"strings"
)

// Request ...
type Request struct {
	function   Function
	symbol     string
	interval   Interval
	outputSize OutputSize
	dataType   DataType
	apiKey     string
}

// NewRequest ...
func NewRequest(function Function, symbol string, interval Interval, outputSize OutputSize,
	dataType DataType, apiKey string) *Request {
	return &Request{
		function:   function,
		symbol:     symbol,
		interval:   interval,
		outputSize: outputSize,
		dataType:   dataType,
		apiKey:     apiKey,
	}
}

// GetURL : generate URL based on input
func (req *Request) GetURL() string {

	var q []string

	q = append(q, fmt.Sprintf("function=%s", req.function))
	q = append(q, fmt.Sprintf("apikey=%s", req.apiKey))
	q = append(q, fmt.Sprintf("symbol=%s", req.symbol))
	q = append(q, fmt.Sprintf("datatype=%s", req.dataType))

	switch req.function {
	case TimeSeriesIntraday:
		q = append(q, fmt.Sprintf("interval=%s", req.interval))
	case TimeSeriesDaily, TimeSeriesDailyAdjusted:
		q = append(q, fmt.Sprintf("outputsize=%s", req.outputSize))
	case TimeSeriesWeekly, TimeSeriesWeeklyAdjusted, TimeSeriesMonthly, TimeSeriesMonthlyAdjusted:
	}

	return fmt.Sprintf("%s/query?%s", ALPHAVANTAGE, strings.Join(q, "&"))
}
