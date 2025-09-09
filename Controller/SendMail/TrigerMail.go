package TrigerMailController

import (
	"net/http"
	TrigerMailModel "nivasBrandRegistrationBackend/Model/SendMail"
	TrigerMailService "nivasBrandRegistrationBackend/Service/SendMail"

	"github.com/gin-gonic/gin"
)

func SendMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqVal TrigerMailModel.SendMailRequest
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}
		resVal := TrigerMailService.SendMail(reqVal)
		response := gin.H{
			"status":  resVal.Status,
			"message": resVal.Message,
		}
		c.JSON(http.StatusOK, gin.H{
			"data": response,
		})
	}
}
