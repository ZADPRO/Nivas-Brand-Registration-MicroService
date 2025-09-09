package TrigerMailRoutes

import (
	TrigerMailController "nivasBrandRegistrationBackend/Controller/SendMail"

	"github.com/gin-gonic/gin"
)

func SendMail(router *gin.Engine) {
	route := router.Group("/brand-on-boding/api/v1/contact")
	route.POST("/SendMail", TrigerMailController.SendMail())
}
