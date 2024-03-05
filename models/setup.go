package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"marketplace-finder/config"
)

var DB *gorm.DB

func ConnectDatabase() {

	dbPath := ""

	if config.CfgValues.DBType == "sqlite" {
		dbPath = "./database.db"
	}

	if config.CfgValues.DBType == "postgres" {
		dbPath = "postgres://" + config.CfgValues.DBUser + ":" + config.CfgValues.DBPassword + "@localhost:" + config.CfgValues.DBPort + "/" + config.CfgValues.DBName
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Target{})
	database.AutoMigrate(&Promocode{})

	DB = database
}
