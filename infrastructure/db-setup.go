package infrastructure

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	server   = "localhost"
	user     = "sa"
	password = "stone-api@123"
	database = "master"
)

func ConnectDatabase() {

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
		server, user, password, database)

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Failed to open database")
	}

	DB = db
}
