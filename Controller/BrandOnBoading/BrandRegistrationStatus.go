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

func GetBrandRegistrationStatusData() gin.HandlerFunc {
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
		var reqVal brandRegistrationModel.GetBrandRegistrationStatusReq
		reqVal.ApplicationId = fmt.Sprintf("%v", applicationId)

		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := brandRegistrationService.GetBrandStatusData(dbConnt, reqVal)
		// response := gin.H{
		// 	"status":    resVal.Status,
		// 	"message":   resVal.Message,
		// 	"brandData": resVal.BrandData,
		// }

		// c.JSON(http.StatusOK, gin.H{
		// 	"data": response,
		// })

		response := gin.H{
			"status":    resVal.Status,
			"message":   resVal.Message,
			"brandData": resVal.BrandData,
		}

		c.JSON(http.StatusOK, gin.H{
			"data": hashapi.Encrypt(response, false, ""),
		})
	}
}
