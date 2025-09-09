package signUpRoutes

import (
	signUpController "nivasBrandRegistrationBackend/Controller/Authentation"

	"github.com/gin-gonic/gin"
)

func SignUpDataRoutes(router *gin.Engine) {
	route := router.Group("/brand-on-boding/api/v1/app/signUp")
	route.POST("/googleLogin", signUpController.SignUpController())
	route.GET("/countryCode", signUpController.CountryCodeController())
	route.POST("/mobileNumberValidation", signUpController.MobileNumberValidationController())
}
