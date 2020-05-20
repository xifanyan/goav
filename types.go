package goav

// ALPHAVANTAGE - root URL for alphavantage
const (
	ALPHAVANTAGE = "https://www.alphavantage.co"
)

// TimeSeriesFunction ...
type TimeSeriesFunction uint8

const (
	// TimeSeriesDaily ...
	TimeSeriesDaily TimeSeriesFunction = iota
	// TimeSeriesDailyAdjusted ...
	TimeSeriesDailyAdjusted
	// TimeSeriesWeekly ...
	TimeSeriesWeekly
	// TimeSeriesWeeklyAdjusted ...
	TimeSeriesWeeklyAdjusted
	// TimeSeriesMonthly ...
	TimeSeriesMonthly
	// TimeSeriesMonthlyAdjusted ...
	TimeSeriesMonthlyAdjusted
	// TimeSeriesIntraday ...
	TimeSeriesIntraday
)

func (f TimeSeriesFunction) String() string {
	return [...]string{
		"TIME_SERIES_DAILY",
		"TIME_SERIES_DAILY_ADJUSTED",
		"TIME_SERIES_WEEKLY",
		"TIME_SERIES_WEEKLY_ADJUSTED",
		"TIME_SERIES_MONTHLY",
		"TIME_SERIES_MONTHLY_ADJUSTED",
		"TIME_SERIES_INTRADAY",
		"EMA",
		"SMA",
	}[f]
}

// Interval ...
type Interval uint8

const (
	// NA ...
	NA Interval = iota
	// OneMinute ...
	OneMinute
	// FiveMinute ...
	FiveMinute
	// FifteenMinute ...
	FifteenMinute
	// ThirtyMinute ...
	ThirtyMinute
	// SixtyMinute ...
	SixtyMinute
	// Daily ...
	Daily
	// Weekly ...
	Weekly
	// Monthly ...
	Monthly
)

func (v Interval) String() string {
	return [...]string{
		"na",
		"1min",
		"5min",
		"15min",
		"30min",
		"60min",
		"daily",
		"weekly",
		"monthly",
	}[v]
}

// DataType ...
type DataType uint8

const (
	// CSV ...
	CSV DataType = iota
	// JSON ...
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
	// COMPACT ...
	COMPACT OutputSize = iota
	// FULL ...
	FULL
)

func (size OutputSize) String() string {
	return [...]string{
		"compact",
		"full",
	}[size]
}
