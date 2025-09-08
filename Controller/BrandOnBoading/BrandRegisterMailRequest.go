package brandRegistrationController

import (
	"net/http"
	db "nivasBrandRegistrationBackend/DB"
	accesstoken "nivasBrandRegistrationBackend/Helper/AccessToken"
	hashapi "nivasBrandRegistrationBackend/Helper/HashAPI"
	brandRegistrationModel "nivasBrandRegistrationBackend/Model/BrandOnBoading"
	brandRegistrationService "nivasBrandRegistrationBackend/Service/BrandOnBoading"

	"github.com/gin-gonic/gin"
)

func GenerateBrandRegisterUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqVal brandRegistrationModel.BrandRegistrationUrlRequest
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}
		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()
		resVal := brandRegistrationService.GenerateBrandRegisterUrl(dbConnt, reqVal)
		token := accesstoken.CreateToken()
		response := gin.H{
			"status":  resVal.Status,
			"message": resVal.Message,
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  hashapi.Encrypt(response, true, token),
			"token": token,
		})
	}
}
