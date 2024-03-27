package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	var err error
	dsn := "Nikhilsingh:password@tcp(127.0.0.1:3306)/NIKHIL?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect with Database")
	}

	DB.AutoMigrate(&BookInventory{}, &Library{}, &IssueRegistery{}, &RequestEvents{}, &Users{})

}
