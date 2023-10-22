package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	mylog "github.com/losevs/people-api/logger"
	"github.com/losevs/people-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func init() {
	// err := godotenv.Load("C:\\Users\\Owner\\Documents\\GoLang\\people-api\\.env")
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		mylog.Logg.Debug("Failed to connect to DB.")
		log.Fatalln(err)
	}
	mylog.Logg.Info("Connected to DB")

	db.Logger = logger.Default.LogMode(logger.Info)

	//aitoMigrate
	err = db.AutoMigrate(&models.PersonResponse{})
	if err != nil {
		mylog.Logg.Debug("migrating error")
		log.Fatalln(err)
	}

	DB = Dbinstance{
		Db: db,
	}
}
