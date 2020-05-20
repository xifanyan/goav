package goav

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DbConfig ...
type DbConfig struct {
	Dialect       string
	ConnectionStr string
}

// Database ...
type Database struct {
	DbConfig
	inst *gorm.DB
}

// NewDatabase ...
func NewDatabase(cfg DbConfig) (*Database, error) {
	var err error

	db := Database{}
	db.DbConfig = cfg

	if db.inst, err = gorm.Open(cfg.Dialect, cfg.ConnectionStr); err != nil {
		return nil, err
	}

	// db.inst.LogMode(true)
	return &db, nil
}

// Init ...
func (db *Database) Init() {
	tbls := []interface{}{&StockRequest{}, &StockTimeSeries{}}
	db.inst.AutoMigrate(tbls...)
}
