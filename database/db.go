package database

import (
	"fmt"
	"submission2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "hudahuda"
	dbname   = "db-go-sql"
	db       *gorm.DB
	err      error
)

func DBInit() *gorm.DB {
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
