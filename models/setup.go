package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	username := os.Getenv("DB_ROOT_PASSWORD")
	password := os.Getenv("DB_ROOT_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dbUrl))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db
}
