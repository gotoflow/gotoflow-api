package internal_database

import (
	"fmt"
	"log"

	config "github.com/gotoflow/core/handlers/config"
	database "github.com/gotoflow/core/handlers/database"
	"gorm.io/gorm"
)

func GetClientDatabase() *gorm.DB {
	config := config.LoadConfig()
	fmt.Println(config.Database.Host)
	db:= database.GetDatabase(config.Database.Driver, config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Database)
	if (db.Error != nil) {
		log.Fatalf("Error on connect DB %s", db.Error)
		panic(db.Error)
	}
	return db
}