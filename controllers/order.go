package controllers

import (
	"fmt"
	"log"
	"net/http"

	"submission2/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (db *InDB) CreateOrder(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)
	id := uuid.New().String()
	order.ID = id
	if err := c.BindJSON(&order); err != nil {
		panic(err)
	}
	db.DB.Create(&order)
	result = gin.H{
		"order": order,
	}
	c.JSON(http.StatusOK, result)
}

func (db *InDB) GetOrder(c *gin.Context) {
	var orders []models.Order
	err := db.DB.Preload("Items").Find(&orders).Error
	if err != nil {
		panic(err)
	}
	result := gin.H{
		"order": orders,
	}
	c.JSON(http.StatusOK, result)
}

func (db *InDB) GetOrderById(c *gin.Context) {
	params := c.Param("id")
	var (
		result     gin.H
		codeResult int
		order      models.Order
	)
	err := db.DB.Preload("Items").First(&order, "id = ?", params).Error
	fmt.Println(err)
	if err != nil {
		result = gin.H{
			"message": "Cannot found ID in database",
		}
		codeResult = http.StatusNotFound
	} else {
		db.DB.First(&models.Order{}, "id = ?", params)
		result = gin.H{
			"order": order,
		}
		codeResult = 202
	}

	c.JSON(codeResult, result)
}

func (db *InDB) UpdateOrder(c *gin.Context) {
	params := c.Param("id")
	log.Println("Update Order", params)
	//notUpdatedOrder := models.Order{}
	var (
		result       gin.H
		UpdatedOrder models.Order
	)
	// db.First(&UpdatedOrder, "id = ?", params)
	UpdatedOrder.ID = params
	if err := c.BindJSON(&UpdatedOrder); err != nil {
		panic(err)
	}

	// err := db.Model(&notUpdatedOrder).Where("id = ?", params).Updates(&UpdatedOrder).Error()
	// if err != nil {
	// 	fmt.Printf("Error updating order")
	// }
	db.DB.Where("order_id= ?", params).Delete(&models.Item{})
	db.DB.Save(&UpdatedOrder)
	result = gin.H{
		"order": UpdatedOrder,
	}

	c.JSON(http.StatusOK, result)
}

func (db *InDB) DeleteOrder(c *gin.Context) {
	params := c.Param("id")
	var (
		result     gin.H
		codeResult int
	)
	err := db.DB.First(&models.Order{}, "id = ?", params).Error
	fmt.Println(err)
	if err != nil {
		result = gin.H{
			"message": "Cannot found ID in database",
		}
		codeResult = http.StatusNotFound
	} else {
		db.DB.Where("order_id= ?", params).Delete(&models.Item{})
		db.DB.Where("id= ?", params).Delete(&models.Order{})

		result = gin.H{
			"message": "Successfully deleted the order",
		}
		codeResult = 202
	}

	c.JSON(codeResult, result)
}
