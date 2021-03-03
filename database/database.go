package database

import (
	"log"

	"jwt/models"
	"jwt/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var dberror error

// OpenDb established connection to database
func OpenDb() *gorm.DB {
	config, configErr := utils.LoadConfig(".")
	if configErr != nil {
		log.Fatal(configErr)
		panic(configErr)
	}

	Db, dberror = gorm.Open(postgres.Open(config.ConnString), &gorm.Config{})
	if dberror != nil {
		log.Fatal(dberror)
		//panic(dberror)
	}

	// Migrations
	Db.AutoMigrate(models.User{})

	return Db
}

// Closedb closes the database connection
func Closedb(db *gorm.DB) {
	sqldb, _ := db.DB()
	sqldb.Close()
}
