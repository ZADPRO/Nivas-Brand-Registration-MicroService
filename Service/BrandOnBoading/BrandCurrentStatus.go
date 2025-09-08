package brandRegistrationService

import (
	"fmt"
	logger "nivasBrandRegistrationBackend/Helper/Logger"
	brandRegistrationModel "nivasBrandRegistrationBackend/Model/BrandOnBoading"
	brandRegistrationQuery "nivasBrandRegistrationBackend/Query/BrandOnBoading"

	"gorm.io/gorm"
)

func GetBrandCurrentStatus(db *gorm.DB, reqVal brandRegistrationModel.GetBrandCurrentStatusReq) brandRegistrationModel.GetBrandCurrentStatusRes {
	log := logger.InitLogger()

	var brandDetails brandRegistrationModel.GetBrandStatusFromDbRes

	err := db.Raw(brandRegistrationQuery.GetBrandCurrentStatus, reqVal.ApplicationId).
		Scan(&brandDetails).Error
	if err != nil {
		log.Error("Error in getting the Brand Register Form Data: " + err.Error())
		return brandRegistrationModel.GetBrandCurrentStatusRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}
	fmt.Println("Db Data : ", brandDetails)

	return brandRegistrationModel.GetBrandCurrentStatusRes{
		Status:    true,
		Message:   "Brand registration data fetched successfully",
		AppStatus: brandDetails.ApplicationStatus,
	}
}
