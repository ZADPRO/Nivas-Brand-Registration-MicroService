package brandDocumentRoutes

import (
	registerFormDocumentController "nivasBrandRegistrationBackend/Controller/BrandDocument"

	"github.com/gin-gonic/gin"
)

func RegisterFormDocumentUploadUrl(router *gin.Engine) {
	route := router.Group("/brand-on-boding/api/v1/brand/registerFormDocument")
	route.POST("/generateUploadUrl", registerFormDocumentController.BrandDocumentUploadUrl())
	route.POST("/DeleteDocument", registerFormDocumentController.DocumentDelete())
}
