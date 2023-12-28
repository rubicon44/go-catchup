package config

import (
	"app/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=sample_db user=user dbname=sample sslmode=disable password=password")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Task{})

	return db
}
