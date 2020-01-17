package goav

// ALPHAVANTAGE - root URL for alphavantage
const (
	ALPHAVANTAGE = "https://www.alphavantage.co"
)

// Function ...
type Function uint8

const (
	TimeSeriesDaily Function = iota
	TimeSeriesDailyAdjusted
	TimeSeriesWeekly
	TimeSeriesWeeklyAdjusted
	TimeSeriesMonthly
	TimeSeriesMonthlyAdjusted
	TimeSeriesIntraday
)

func (f Function) String() string {
	return [...]string{
		"TIME_SERIES_DAILY",
		"TIME_SERIES_DAILY_ADJUSTED",
		"TIME_SERIES_WEEKLY",
		"TIME_SERIES_WEEKLY_ADJUSTED",
		"TIME_SERIES_MONTHLY",
		"TIME_SERIES_MONTHLY_ADJUSTED",
		"TIME_SERIES_INTRADAY",
	}[f]
}

// Interval ...
type Interval uint8

const (
	IntervalOneMinute Interval = iota
	IntervalFiveMinute
	IntervalFifteenMinute
	IntervalThirtyMinute
	IntervalSixtyMinute
	IntervalDaily
	IntervalWeekly
	IntervalMonthly
	IntervalNA
)

func (v Interval) String() string {
	return [...]string{
		"1min",
		"5min",
		"15min",
		"30min",
		"60min",
		"1day",
		"1mon",
		"na",
	}[v]
}

// DataType ...
type DataType uint8

const (
	CSV DataType = iota
	JSON
)

func (dtype DataType) String() string {
	return [...]string{
		"csv",
		"json",
	}[dtype]
}

// OutputSize ...
type OutputSize uint8

const (
	COMPACT OutputSize = iota
	FULL
)

func (size OutputSize) String() string {
	return [...]string{
		"compact",
		"full",
	}[size]
}