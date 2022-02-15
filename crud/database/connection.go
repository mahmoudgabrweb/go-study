package database

import (
	"github.com/jinzhu/gorm"
	"log"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("DB is connected successfully.")
	return nil
}


