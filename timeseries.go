package goav

type TimeSeries struct {
	req requestor
	db  *Database
}

func NewTimeSeries(req requestor, db *Database) *TimeSeries {
	return &TimeSeries{
		req: req,
		db:  db,
	}
}
