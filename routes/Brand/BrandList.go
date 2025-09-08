package routes

import (
	controller "nivasBrandRegistrationBackend/Controller/Brand"

	"github.com/gin-gonic/gin"
)

func PendingBrandList(router *gin.Engine) {
	route := router.Group("brand-on-boding/api/v1/brand")
	route.POST("/pendingList", controller.BrandPendingList())
	route.POST("/approvedList", controller.BrandApprovedList())
	route.POST("/rejectedList", controller.BrandRejectedList())
	route.POST("/brandInformation", controller.BrandInformation())
	route.POST("/updateBrandStatus", controller.BrandStatusUpdate())
}
