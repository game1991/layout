package store

import (
	"helloworld/pkg/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config ...
type Config struct {
	DSN string
}

// NewMySQL ...
func NewMySQL(config Config) (db *gorm.DB, cleanup func(), err error) {
	db, err = gorm.Open(
		mysql.Open(config.DSN),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	cleanup = func() {
		log.Info("closing the data resources")
	}
	return
}
