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
		code   int
		item   models.Item
	)
	id := uuid.New().String()
	order.ID = id
	if err := c.BindJSON(&order); err != nil {
		panic(err)
	}
	fmt.Println(order.Items, item)
	if len(order.Items) <= 0 {
		code = 409
		result = gin.H{
			"message": "Cannot put order, item is empty",
		}
	} else if len(order.Items) > 0 {
		db.DB.Create(&order)
		code = 200
		result = gin.H{
			"message": "Successfully created order",
			"order":   order,
		}
	} else {
		result = gin.H{
			"message": "Something went wrong, please try again the input",
		}
	}
	c.JSON(code, result)
}

func (db *InDB) GetOrder(c *gin.Context) {
	var (
		orders []models.Order
		code   int
		result gin.H
	)
	err := db.DB.Preload("Items").Find(&orders).Error
	if err != nil {
		result = gin.H{
			"message": "Something went wrong, please try again",
		}
		panic(err)
	} else {
		code = http.StatusOK
		result = gin.H{
			"order": orders,
		}
		c.JSON(code, result)
	}

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
	var (
		result       gin.H
		UpdatedOrder models.Order
		code         int
	)
	UpdatedOrder.ID = params
	if err := c.BindJSON(&UpdatedOrder); err != nil {
		panic(err)
	}

	if err := db.DB.First(&models.Order{}, "id = ?", params).Error; err != nil {
		result = gin.H{
			"message": "Cannot found ID in database",
		}
		code = http.StatusNotFound
	} else if len(UpdatedOrder.Items) <= 0 {
		code = 409
		result = gin.H{
			"message": "Cannot edit order, item is empty",
		}
	} else if len(UpdatedOrder.Items) > 0 {
		db.DB.Where("order_id= ?", params).Delete(&models.Item{})
		db.DB.Save(&UpdatedOrder)
		result = gin.H{
			"message": "Successfully updated order",
			"order":   UpdatedOrder,
		}
	} else {
		result = gin.H{
			"message": "Something went wrong, please try again the input",
		}
	}

	c.JSON(code, result)
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
