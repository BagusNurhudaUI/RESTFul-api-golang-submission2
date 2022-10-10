package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var message string
	c.JSON(http.StatusOK, gin.H{
		message: "Successfully get into Index",
	})
}
