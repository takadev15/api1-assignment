package database

import (
	"fmt"
	"log"
	// "os"

	"github.com/takadev15/go-restapi/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	userName = "taka"
	dbName   = "orders_by"
	dbPass   = "hacktiv8pass"
	dbPort   = "5432"
	db       *gorm.DB
	err      error
)

func DBinit() {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", userName, dbPass, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Databases Error", err.Error())
	}
	log.Printf("Databases Connected")
	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
