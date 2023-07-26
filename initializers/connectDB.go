package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error

	// A string used to describe a connection to data source
	dsn := fmt.Sprintf("host=%s, user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=EST", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	// Uses GORM to objectify db
	// postgres.Open(dsn)	- Driver that connects using DSN
	// &gorm.config{} 		- Empty configuration
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to conect to the Database")
	}

	fmt.Println("? Connected Successfully to the Database")
}
