package database

import (
	"log"
	"os"

	"github.com/atageldi194229/e-deal-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance
var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=atasan password=atasan dbname=e_deal port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	DB.AutoMigrate(&models.Category{})

	Database = DbInstance{
		Db: DB,
	}
}

func Automigrate() {
	// log.Println("Running Migrations")
	//
	// DB.AutoMigrate(&models.Category{})
}
