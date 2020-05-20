package goav

import (
	"testing"
)

func TestURLs(t *testing.T) {
	req := NewStockRequest(TimeSeriesDailyAdjusted, "MSFT", NA, COMPACT, CSV, "demo")

	if req.getURL() != "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&apikey=demo&symbol=MSFT&datatype=csv&outputsize=compact" {
		t.Errorf("wrong URL %s", req.getURL())
	}
}
