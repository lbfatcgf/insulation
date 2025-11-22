package database

import (
	"time"

	"insulation/server/base/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize() {
	if db != nil {
		return
	}
	dns := config.Global().DataSource.DataBase.DSN
	opendb, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	opendb.Session(&gorm.Session{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	db = opendb
}

func GetDB() *gorm.DB {
	return db
}
