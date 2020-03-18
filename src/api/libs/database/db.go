package database

import (
	"api/utils/console"
	"config"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var once sync.Once

var instance *gorm.DB

// Connect to the DATABASE
func Connect() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		console.Warning("Connecting Database")
		instance, err = gorm.Open(config.DBDRIVER, config.DBURL)
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		instance.DB().SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		instance.DB().SetMaxOpenConns(100)
	})
	if err != nil {
		return nil, err
	}
	console.Warning("Connected Database")
	if config.DBDEBUG == true {
		instance = instance.Debug()
	}
	return instance, nil
}
