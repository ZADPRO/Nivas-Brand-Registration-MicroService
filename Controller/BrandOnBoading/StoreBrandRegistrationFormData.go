package brandRegistrationController

import (
	"fmt"
	"net/http"
	db "nivasBrandRegistrationBackend/DB"
	hashapi "nivasBrandRegistrationBackend/Helper/HashAPI"
	brandRegistrationModel "nivasBrandRegistrationBackend/Model/BrandOnBoading"
	brandRegistrationService "nivasBrandRegistrationBackend/Service/BrandOnBoading"

	"github.com/gin-gonic/gin"
)

func StoreBrandRegistrationForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		applicationId, applicationIdExit := c.Get("applicationId")

		if !applicationIdExit {
			// Handle error: ID is missing from context (e.g., middleware didn't set it)
			c.JSON(http.StatusUnauthorized, gin.H{ // Or StatusInternalServerError depending on why it's missing
				"status":  false,
				"message": "Application Id Is Not Presented.",
			})
			return
		}
		var reqVal brandRegistrationModel.StoreBrandRegisterFormDataReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}
		reqVal.ApplicationId = fmt.Sprintf("%v", applicationId)

		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := brandRegistrationService.StoreBrandDetails(dbConnt, reqVal)
		// response := gin.H{
		// 	"status":            resVal.Status,
		// 	"message":           resVal.Message,
		// 	"applicationStatus": resVal.ApplicationStatus,
		// }

		// c.JSON(http.StatusOK, gin.H{
		// 	"data": response,
		// })

		response := gin.H{
			"status":            resVal.Status,
			"message":           resVal.Message,
			"applicationStatus": resVal.ApplicationStatus,
		}

		c.JSON(http.StatusOK, gin.H{
			"data": hashapi.Encrypt(response, false, ""),
		})

	}
}
