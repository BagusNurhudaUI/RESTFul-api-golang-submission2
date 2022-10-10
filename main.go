package main

import (
	"fmt"
	"submission2/controllers"
	"submission2/database"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.DBInit()
	InDB := controllers.New(db)
	fmt.Println("Starting server")
	r := gin.Default()

	r.GET("/", controllers.Index)
	r.POST("/orders", InDB.CreateOrder)
	r.GET("/orders", InDB.GetOrder)
	r.GET("/orders/:id", InDB.GetOrderById)
	r.PUT("/orders/:id", InDB.UpdateOrder)
	r.DELETE("/orders/:id", InDB.DeleteOrder)
	r.Run("127.0.0.1:3000")
}
