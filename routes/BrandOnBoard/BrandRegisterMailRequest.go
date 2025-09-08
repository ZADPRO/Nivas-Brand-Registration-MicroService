package brandOnBodingRoutes

import (
	brandRegistrationController "nivasBrandRegistrationBackend/Controller/BrandOnBoading"
	accesstoken "nivasBrandRegistrationBackend/Helper/AccessToken"

	"github.com/gin-gonic/gin"
)

func BrandOnBoading(router *gin.Engine) {
	route := router.Group("/brand-on-boding/api/v1/brand")
	route.POST("/BrandRegisterMailRequest", brandRegistrationController.GenerateBrandRegisterUrl())
	route.GET("/getStatus", accesstoken.JWTMiddleware(), brandRegistrationController.GetBrandCurrentStatus())
	route.GET("/getFormData", accesstoken.JWTMiddleware(), brandRegistrationController.GetBrandRegistrationData())
	route.GET("/getRegistrationStatus", accesstoken.JWTMiddleware(), brandRegistrationController.GetBrandRegistrationStatusData())
	route.POST("/storeData", accesstoken.JWTMiddleware(), brandRegistrationController.StoreBrandRegistrationForm())

}
