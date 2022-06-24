package db

import (
	"log"
	"os"

	"github.com/adilsonmenechini/simplesrestapi/api/utils/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

//ConnectPSQL is a function for the connection of the db
func ConnectPSQL() *gorm.DB {
	var err error

	config.Enviroment()

	dsn := os.Getenv("psql")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), //logger blocked
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}
	log.Println("connected postgres")
	migrate(db)
	return db
}
