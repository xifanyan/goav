package goav

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DbConfig struct {
	Dialect       string
	ConnectionStr string
}

type Database struct {
	DbConfig
	db *gorm.DB
}

func NewDatabase(cfg DbConfig) (*Database, error) {
	var err error

	database := Database{}
	database.DbConfig = cfg

	if database.db, err = gorm.Open(cfg.Dialect, cfg.ConnectionStr); err != nil {
		return nil, err
	}

	database.db.LogMode(true)

	return &database, nil
}

func (database *Database) Init() {

	database.db.AutoMigrate(&StockRequest{})

}
