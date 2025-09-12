package brandRegistrationModel

type InvitedBrandList struct {
	BrandId       int    `json:"brandId" gorm:"column:refApplicationId"`
	BrandCustId   string `json:"brandCustId" gorm:"column:refApplicationCustId"`
	BrandName     string `json:"brandName" gorm:"column:refBrandName"`
	ContactPerson string `json:"brandContactPerson" gorm:"column:refBrandContactPerson"`
	ContactEmail  string `json:"mailId" gorm:"column:refMailId"`
	DateOfInvite  string `json:"inviteDate" gorm:"column:refCreateAT"`
	Status        string `json:"status" gorm:"column:refStatus"`
}

type InvitedBrandListRes struct {
	BrandList []InvitedBrandList `json:"brandList"`
	Message   string             `json:"message"`
	Status    bool               `json:"status"`
}

type ReInviteBrandRequest struct {
	BrandId int `json:"brandId"`
}

type ReInviteBrandResponce struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type InviteData struct {
	ContactPerson   string
	BrandName       string
	RegistrationURL string
	BrandMailId     string
	BrandMobile     string
}
