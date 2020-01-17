package goav

import (
	"testing"
)

func TestURLs(t *testing.T) {
	req := NewRequest(TimeSeriesDailyAdjusted, "MSFT", IntervalNA, COMPACT, CSV, "demo")

	if req.GetURL() != "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&apikey=demo&symbol=MSFT&datatype=csv&outputsize=compact" {
		t.Errorf("wrong URL %s", req.GetURL())
	}
}
