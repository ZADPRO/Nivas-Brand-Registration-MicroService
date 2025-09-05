package model

type PendingBrandListReq struct {
	CategoryId []int `json:"categoryId"`
}

type BrandListQueryRes struct {
	ApplicationId     int    `json:"refApplicationId" gorm:"column:refApplicationId"`
	BrandId           string `json:"refBrandId" gorm:"column:refBrandId"`
	ApplicationCustId string `json:"refApplicationCustId" gorm:"column:refApplicationCustId"`
	BrandName         string `json:"refBrandName" gorm:"column:refBrandName"`
	CategoryName      string `json:"refBrandCategoryName" gorm:"column:refBrandCategoryName"`
	ContactPerson     string `json:"refBrandContactPerson" gorm:"column:refBrandContactPerson"`
	IfWareHouse       bool   `json:"refIfWareHouse" gorm:"column:refIfWareHouse"`
	JoinAt            string `json:"refCreateAT" gorm:"column:refCreateAT"`
	UpdateAt          string `json:"refUpdateAt" gorm:"column:refUpdateAt"`
	ApplicationStatus string `json:"refApplicationStatus" gorm:"column:refApplicationStatus"`
	BrandLogo         string `json:"refLogo" gorm:"column:refLogo"`
}

type PendingBrandList struct {
	BrandId           int    `json:"brandId"`
	BrandCustId       string `json:"brandCustId"`
	BrandName         string `json:"brandName"`
	Category          string `json:"category"`
	ContactPerson     string `json:"contactPerson"`
	Warehouse         string `json:"warehouse"`
	DateOfApplication string `json:"dateOfApplication"`
	DateOfUpdate      string `json:"dateOfUpdate"`
	Status            string `json:"status"`
	Logo              string `json:"logo`
}

type PendingBrandListRes struct {
	Message   string             `json:"message"`
	Status    bool               `json:"status"`
	BrandList []PendingBrandList `json:"brandList"`
}

type BrandInformationReq struct {
	BrandId int `json:"brandId"`
}

type GetBrandInformationFromDb struct {
	BrandName           string `json:"refBrandName" gorm:"column:refBrandName"`
	ProductCategory     int    `json:"refProductCatageoryId" gorm:"column:refProductCatageoryId"`
	BrandLogoPath       string `json:"refLogo" gorm:"column:refLogo"`
	BrandDescription    string `json:"refBrandDesciption" gorm:"column:refBrandDesciption"`
	WebsiteURL          string `json:"refWebsiteUrl" gorm:"column:refWebsiteUrl"`
	Instragram          string `json:"refInstaUrl" gorm:"column:refInstaUrl"`
	ContactPerson       string `json:"refBrandContactPerson" gorm:"column:refBrandContactPerson"`
	Designation         string `json:"refDesignation" gorm:"column:refDesignation"`
	PhoneNumber         string `json:"refBrandMobile" gorm:"column:refBrandMobile"`
	Email               string `json:"refBrandEmail" gorm:"column:refBrandEmail"`
	Address             string `json:"refAddress" gorm:"column:refAddress"`
	City                string `json:"refCity" gorm:"column:refCity"`
	ZipCode             string `json:"refZipCode" gorm:"column:refZipCode"`
	State               string `json:"refState" gorm:"column:refState"`
	ProofDocument       string `json:"refAddressProf" gorm:"column:refAddressProf"`
	GstinNumber         string `json:"refGstin" gorm:"column:refGstin"`
	CinNumber           string `json:"refCin" gorm:"column:refCin"`
	GstDocumant         string `json:"refGstinPath" gorm:"column:refGstinPath"`
	PanDocument         string `json:"refPanCars" gorm:"column:refPanCars"`
	WareHouse           bool   `json:"refIfWareHouse" gorm:"column:refIfWareHouse"`
	WareHouseAddress    string `json:"wareHouseAddress" gorm:"column:wareHouseAddress"`
	WareHouseCity       string `json:"wareHouseCity" gorm:"column:wareHouseCity"`
	WareHouseDistrict   string `json:"wareHouseDistrict" gorm:"column:wareHouseDistrict"`
	WareHouseZipCode    string `json:"wareHouseZipCode" gorm:"column:wareHouseZipCode"`
	WareHouseState      string `json:"wareHouseState" gorm:"column:wareHouseState"`
	SaveDraft           bool   `json:"refSaveDraft" gorm:"column:refSaveDraft"`
	ApplicationStatusId int    `json:"refApplicationStatusId" gorm:"column:refApplicationStatusId"`
	ApplicationStatus   string `json:"refApplicationStatus" gorm:"column:refApplicationStatus"`
	BrandLogoName       string `json:"refLogoName" gorm:"column:refLogoName"`
	ProffDocumentName   string `json:"refAddressProfName" gorm:"column:refAddressProfName"`
	GstDocumentName     string `json:"refGstinName" gorm:"column:refGstinName"`
	PanDocumantName     string `json:"refPanCarsName" gorm:"column:refPanCarsName"`
}

type DocumentLink struct {
	Url         string `json:"url"`
	DownloadUrl string `json:"downloadUrl"`
	FileSize    string `json:"fileSize"`
}
type GetBrandRegisterDataResponse struct {
	BrandInformation struct {
		BrandName        string `json:"brandName"`
		ProductCategory  int    `json:"productCategory"`
		BrandLogoPath    string `json:"brandLogoPath"`
		BrandDescription string `json:"brandDescription"`
		WebsiteURL       string `json:"websiteURL"`
		Instragram       string `json:"instragram"`
	} `json:"brandInformation"`
	ContactInformation struct {
		ContactPerson string `json:"contactPerson"`
		Designation   string `json:"designation"`
		PhoneNumber   string `json:"phoneNumber"`
		Email         string `json:"email"`
		Address       string `json:"address"`
		City          string `json:"city"`
		ZipCode       string `json:"zipCode"`
		State         string `json:"state"`
		ProofDocument string `json:"proofDocument"`
	} `json:"contactInformation"`
	TaxInformation struct {
		GstinNumber string `json:"gstinNumber"`
		CinNumber   string `json:"cinNumber"`
		GstDocumant string `json:"gstDocumant"`
		PanDocument string `json:"panDocument"`
	} `json:"taxInformation"`
	WareHouseInfo struct {
		WareHouse         bool   `json:"wareHouse"`
		WareHouseAddress  string `json:"wareHouseAddress"`
		WareHouseCity     string `json:"wareHouseCity"`
		WareHouseDistrict string `json:"wareHouseDistrict"`
		WareHouseZipCode  string `json:"wareHouseZipCode"`
		WareHouseState    string `json:"wareHouseState"`
	} `json:"wareHouseInfo"`
	ApplicationType struct {
		SaveDraft              bool   `json:"saveDraft"`
		RefApplicationStatusId int    `json:"refApplicationStatusId"`
		RefApplicationStatus   string `json:"refApplicationStatus"`
	} `json:"applicationType"`
	Document struct {
		LogoFile     string       `json:"logoFile"`
		AddressProof DocumentLink `json:"addressProof"`
		GstDocument  DocumentLink `json:"gstDocument"`
		PanDocument  DocumentLink `json:"panDocument"`
	} `json:"document"`
}

type GetBrandInformationRes struct {
	BrandData GetBrandRegisterDataResponse `json:"brandData"`
	Status    bool                         `json:"status"`
	Message   string                       `json:"message"`
}

type BrandStatusUpdateReq struct {
	BrandId      int    `json:"brandId"`
	BrandApprove bool   `json:"brandApprove"`
	Reason       string `json:"reason"`
}

type BrandStatusUpdateResFromDb struct {
	ApplicationId string `json:"refApplicationCustId" gorm:"column:refApplicationCustId"`
	BrandEmail    string `json:"refBrandEmail" gorm:"column:refBrandEmail"`
	BrandName     string `json:"refBrandName" gorm:"column:refBrandName"`
	ContactPerson string `json:"refBrandContactPerson" gorm:"refBrandContactPerson"`
}

type BrandStatusUpdateRes struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
