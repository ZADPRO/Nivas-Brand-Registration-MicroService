package controller

import (
	"net/http"
	db "nivasBrandRegistrationBackend/DB"
	accesstoken "nivasBrandRegistrationBackend/Helper/AccessToken"
	hashapi "nivasBrandRegistrationBackend/Helper/HashAPI"
	model "nivasBrandRegistrationBackend/Model/Brand"
	"os"

	// accesstoken "nivasBrandRegistrationBackend/Helper/AccessToken"
	service "nivasBrandRegistrationBackend/Service/Brand"

	"github.com/gin-gonic/gin"
)

func BrandPendingList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqVal model.PendingBrandListReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}

		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := service.BrandPendingList(dbConnt, reqVal)
		token := accesstoken.CreateToken()

		response := gin.H{
			"status":    resVal.Status,
			"message":   resVal.Message,
			"brandList": resVal.BrandList,
		}

		// pick specific env vars
		envdata := gin.H{
			"ENCRYPT_API": os.Getenv("ENCRYPT_API"),
			// add only what you need, don’t dump secrets!
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    hashapi.Encrypt(response, true, token),
			"token":   token,
			"envdata": envdata,
		})
	}
}
func BrandApprovedList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqVal model.PendingBrandListReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}
		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()
		resVal := service.BrandApprovedList(dbConnt, reqVal)
		token := accesstoken.CreateToken()
		response := gin.H{
			"status":    resVal.Status,
			"message":   resVal.Message,
			"brandList": resVal.BrandList,
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  hashapi.Encrypt(response, true, token),
			"token": token,
		})
	}
}
func BrandRejectedList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqVal model.PendingBrandListReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}
		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()
		resVal := service.BrandRejectedList(dbConnt, reqVal)
		token := accesstoken.CreateToken()
		response := gin.H{
			"status":    resVal.Status,
			"message":   resVal.Message,
			"brandList": resVal.BrandList,
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  hashapi.Encrypt(response, true, token),
			"token": token,
		})
	}
}
func BrandInformation() gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqVal model.BrandInformationReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}
		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()
		resVal := service.BrandInformation(dbConnt, reqVal)
		token := accesstoken.CreateToken()
		response := gin.H{
			"status":    resVal.Status,
			"message":   resVal.Message,
			"brandList": resVal.BrandData,
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  hashapi.Encrypt(response, true, token),
			"token": token,
		})
	}
}
func BrandStatusUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqVal model.BrandStatusUpdateReq
		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again " + err.Error(),
			})
			return
		}
		dbConnt, sqlDB := db.InitDB()
		defer sqlDB.Close()
		resVal := service.BrandStatusUpdate(dbConnt, reqVal)
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
