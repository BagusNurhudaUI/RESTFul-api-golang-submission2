package main

import (
	"fmt"
	"log"
	"os"
	"submission2/controllers"
	"submission2/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	db := database.DBInit()
	InDB := controllers.New(db)
	fmt.Println("Starting server")
	r := gin.Default()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	r.GET("/", controllers.Index)
	r.POST("/orders", InDB.CreateOrder)
	r.GET("/orders", InDB.GetOrder)
	r.GET("/orders/:id", InDB.GetOrderById)
	r.PUT("/orders/:id", InDB.UpdateOrder)
	r.DELETE("/orders/:id", InDB.DeleteOrder)
	r.Run(PORT)
}
