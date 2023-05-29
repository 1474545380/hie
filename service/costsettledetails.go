package service

import (
	"github.com/gin-gonic/gin"
	"hie/main/handler"
	"net/http"
)

func GetallNotSettlement(c *gin.Context) {
	cs, err := handler.GetallNotSettlement()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": cs,
	})
}
