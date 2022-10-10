package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"submission2/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	portdb, err := strconv.Atoi(os.Getenv("portdb"))
	var (
		host     = os.Getenv("host")
		port     = portdb
		user     = os.Getenv("user")
		password = os.Getenv("password")
		dbname   = os.Getenv("dbname")
		db       *gorm.DB
	)
	fmt.Println(port)
	config1 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(config1), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connected to databasee")
	db.AutoMigrate(&models.Order{}, &models.Item{})
	// 	return err1.Error()
	// }
	// if err := db.AutoMigrate(models.Order{}, models.Item{}).Error; err != nil {
	// 	fmt.Println("error migrating")
	// }
	return db
}
