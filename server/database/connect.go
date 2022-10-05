package database

import (
	"log"
	"os"

	"github.com/sohhamm/entertainment-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDB() {
	var err error

	db_url := config.Config("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Error connecting to the db 😞 \n", err)
		os.Exit(2)
	}

	log.Println("Connected to the database 🔥")

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal("Connection Pooling failed")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations .... 🚂")

	db.AutoMigrate()

	DB = DbInstance{Db: db}

}
