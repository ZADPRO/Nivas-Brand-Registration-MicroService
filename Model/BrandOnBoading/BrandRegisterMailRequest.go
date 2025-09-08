package brandRegistrationModel

type BrandRegistrationUrlRequest struct {
	MailId       string `json:"mailId" binding :"required"`
	BrandName    string `json:"brandName" binding :"required"`
	CustomerName string `json:"customerName" binding :"required"`
}

type BrandEmailCheckRes struct {
	Status            bool   `json:"status" binding : "required"`
	Message           string `json:"message" binding :"required"`
	ApplicationId     int    `json:"refApplicationId" gorm:"refApplicationId"`
	BrandId           string `json:"refBrandId" gorm:"refBrandId"`
	ApplicationCustId string `json:"refApplicationCustId" gorm:"refApplicationCustId"`
	BrandName         string `json:"refBrandName" gorm:"refBrandName"`
	MailId            string `json:"refMailId" gorm:"refMailId"`
}

type BrandMailCheckQueryResponse struct {
	ApplicationCustId string `json:"refApplicationCustId" gorm:"column:refApplicationCustId"`
	BrandName         string `json:"refBrandName" gorm:"column:refBrandName"`
	MailId            string `json:"refMailId" gorm:"column:refMailId"`
}

type BrandApplicationIdResponse struct {
	ApplicationCustId string `json:"refApplicationCustId" gorm:"column:refApplicationCustId"`
}

type GetBrandCurrentStatusReq struct {
	ApplicationId string `json:"applicationId"`
}

type GetBrandStatusFromDbRes struct {
	ApplicationStatus int `json:"refApplicationStatus" gorm:"column:refApplicationStatus"`
}

type GetBrandCurrentStatusRes struct {
	AppStatus int    `json:"applicationStatus`
	Status    bool   `json:"status"`
	Message   string `json:"message"`
}
