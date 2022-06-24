package db

import (
	"log"
	"os"

	"github.com/adilsonmenechini/simplesrestapi/api/utils/config"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//ConnectPSQL is a function for the connection of the db
func ConnectMySQL() *gorm.DB {
	var err error

	config.Enviroment()

	dsn := os.Getenv("mysql")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), //logger blocked
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}
	log.Println("connected mysql")
	migrate(db)
	return db
}
