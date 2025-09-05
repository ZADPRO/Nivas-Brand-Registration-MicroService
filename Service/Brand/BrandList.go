package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	becrypt "nivasBrandRegistrationBackend/Helper/Becrypt"
	getDate "nivasBrandRegistrationBackend/Helper/DateFormat"
	logger "nivasBrandRegistrationBackend/Helper/Logger"
	mailService "nivasBrandRegistrationBackend/Helper/MailService"
	service "nivasBrandRegistrationBackend/Helper/MinIo"
	model "nivasBrandRegistrationBackend/Model/Brand"
	query "nivasBrandRegistrationBackend/Query/Brand"
	"os"
	"time"

	"gorm.io/gorm"
)

func BrandPendingList(db *gorm.DB, reqVal model.PendingBrandListReq) model.PendingBrandListRes {
	log := logger.InitLogger()
	var PendingBrandList []model.BrandListQueryRes
	err := db.Raw(query.GetBrandListData, 3, reqVal.CategoryId).Scan(&PendingBrandList).Error
	if err != nil {
		log.Error("Error in Getting the Brand list: " + err.Error())
		return model.PendingBrandListRes{
			Status:  false,
			Message: "Failed to create brand application",
		}
	}

	var response []model.PendingBrandList
	for _, b := range PendingBrandList {
		// Format date
		parsed, _ := time.Parse("2006-01-02", b.JoinAt)
		formattedDate := parsed.Format("Jan 02, 2006")
		parsed1, _ := time.Parse("2006-01-02", b.UpdateAt)
		formattedDate1 := parsed1.Format("Jan 02, 2006")
		addViewUrl, err := service.GetFileURL(b.BrandLogo, 30)
		if err != nil {
			fmt.Println("Error generating BrandLogo URL:", err)
		}

		response = append(response, model.PendingBrandList{
			BrandId:           b.ApplicationId,
			BrandCustId:       b.ApplicationCustId,
			BrandName:         b.BrandName,
			Category:          b.CategoryName,
			ContactPerson:     b.ContactPerson,
			Warehouse:         map[bool]string{true: "Yes", false: "No"}[b.IfWareHouse],
			DateOfApplication: formattedDate,
			Status:            b.ApplicationStatus,
			Logo:              addViewUrl,
			DateOfUpdate:      formattedDate1,
		})
	}

	return model.PendingBrandListRes{
		Message:   "Not Approved Pending Brand List fetched successfully",
		Status:    true,
		BrandList: response,
	}
}
func BrandApprovedList(db *gorm.DB, reqVal model.PendingBrandListReq) model.PendingBrandListRes {
	log := logger.InitLogger()
	var PendingBrandList []model.BrandListQueryRes
	err := db.Raw(query.GetBrandListData, 4, reqVal.CategoryId).Scan(&PendingBrandList).Error
	if err != nil {
		log.Error("Error in Getting the Brand list: " + err.Error())
		return model.PendingBrandListRes{
			Status:  false,
			Message: "Failed to create brand application",
		}
	}

	var response []model.PendingBrandList
	for _, b := range PendingBrandList {
		// Format date
		parsed, _ := time.Parse("2006-01-02", b.JoinAt)
		formattedDate := parsed.Format("Jan 02, 2006")
		parsed1, _ := time.Parse("2006-01-02", b.UpdateAt)
		formattedDate1 := parsed1.Format("Jan 02, 2006")
		addViewUrl, err := service.GetFileURL(b.BrandLogo, 30)
		if err != nil {
			fmt.Println("Error generating BrandLogo URL:", err)
		}

		response = append(response, model.PendingBrandList{
			BrandId:           b.ApplicationId,
			BrandCustId:       b.ApplicationCustId,
			BrandName:         b.BrandName,
			Category:          b.CategoryName,
			ContactPerson:     b.ContactPerson,
			Warehouse:         map[bool]string{true: "Yes", false: "No"}[b.IfWareHouse],
			DateOfApplication: formattedDate,
			Status:            b.ApplicationStatus,
			Logo:              addViewUrl,
			DateOfUpdate:      formattedDate1,
		})
	}

	return model.PendingBrandListRes{
		Message:   "Approved Brand List Passed Successfully",
		Status:    true,
		BrandList: response,
	}
}
func BrandRejectedList(db *gorm.DB, reqVal model.PendingBrandListReq) model.PendingBrandListRes {
	log := logger.InitLogger()
	var PendingBrandList []model.BrandListQueryRes
	err := db.Raw(query.GetBrandListData, 5, reqVal.CategoryId).Scan(&PendingBrandList).Error
	if err != nil {
		log.Error("Error in Getting the Brand list: " + err.Error())
		return model.PendingBrandListRes{
			Status:  false,
			Message: "Failed to create brand application",
		}
	}

	var response []model.PendingBrandList
	for _, b := range PendingBrandList {
		// Format date
		parsed, _ := time.Parse("2006-01-02", b.JoinAt)
		formattedDate := parsed.Format("Jan 02, 2006")
		parsed1, _ := time.Parse("2006-01-02", b.UpdateAt)
		formattedDate1 := parsed1.Format("Jan 02, 2006")
		addViewUrl, err := service.GetFileURL(b.BrandLogo, 30)
		if err != nil {
			fmt.Println("Error generating BrandLogo URL:", err)
		}

		response = append(response, model.PendingBrandList{
			BrandId:           b.ApplicationId,
			BrandCustId:       b.ApplicationCustId,
			Category:          b.CategoryName,
			ContactPerson:     b.ContactPerson,
			Warehouse:         map[bool]string{true: "Yes", false: "No"}[b.IfWareHouse],
			DateOfApplication: formattedDate,
			Status:            b.ApplicationStatus,
			Logo:              addViewUrl,
			DateOfUpdate:      formattedDate1,
		})
	}

	return model.PendingBrandListRes{
		Message:   "Rejected Brand List is Passed Successfully",
		Status:    true,
		BrandList: response,
	}
}

func MapBrandDetails(brandDetails model.GetBrandInformationFromDb) model.GetBrandRegisterDataResponse {
	var resp model.GetBrandRegisterDataResponse

	// Brand Information
	resp.BrandInformation.BrandName = brandDetails.BrandName
	resp.BrandInformation.ProductCategory = brandDetails.ProductCategory
	resp.BrandInformation.BrandLogoPath = brandDetails.BrandLogoPath
	resp.BrandInformation.BrandDescription = brandDetails.BrandDescription
	resp.BrandInformation.WebsiteURL = brandDetails.WebsiteURL
	resp.BrandInformation.Instragram = brandDetails.Instragram

	// Contact Information
	resp.ContactInformation.ContactPerson = brandDetails.ContactPerson
	resp.ContactInformation.Designation = brandDetails.Designation
	resp.ContactInformation.PhoneNumber = brandDetails.PhoneNumber
	resp.ContactInformation.Email = brandDetails.Email
	resp.ContactInformation.Address = brandDetails.Address
	resp.ContactInformation.City = brandDetails.City
	resp.ContactInformation.ZipCode = brandDetails.ZipCode
	resp.ContactInformation.State = brandDetails.State
	resp.ContactInformation.ProofDocument = brandDetails.ProofDocument

	// Tax Information
	resp.TaxInformation.GstinNumber = brandDetails.GstinNumber
	resp.TaxInformation.CinNumber = brandDetails.CinNumber
	resp.TaxInformation.GstDocumant = brandDetails.GstDocumant
	resp.TaxInformation.PanDocument = brandDetails.PanDocument

	// Warehouse Information
	resp.WareHouseInfo.WareHouse = brandDetails.WareHouse
	resp.WareHouseInfo.WareHouseAddress = brandDetails.WareHouseAddress
	resp.WareHouseInfo.WareHouseCity = brandDetails.WareHouseCity
	resp.WareHouseInfo.WareHouseDistrict = brandDetails.WareHouseDistrict
	resp.WareHouseInfo.WareHouseZipCode = brandDetails.WareHouseZipCode
	resp.WareHouseInfo.WareHouseState = brandDetails.WareHouseState

	// Application Type
	resp.ApplicationType.SaveDraft = brandDetails.SaveDraft
	resp.ApplicationType.RefApplicationStatusId = brandDetails.ApplicationStatusId
	resp.ApplicationType.RefApplicationStatus = brandDetails.ApplicationStatus

	logoDoc, err := service.GetFileURL(brandDetails.BrandLogoPath, 30)
	if err != nil {
		fmt.Println("Error generating BrandLogo URL:", err)
	}

	resp.Document.LogoFile = logoDoc
	addViewUrl, _ := service.GetFileURL(brandDetails.ProofDocument, 30)
	addDownloadUrl, _ := service.CreateDownloadURL(brandDetails.ProofDocument, brandDetails.ProffDocumentName, 30)
	addressDocSize, _ := service.GetFileInfo(brandDetails.ProofDocument)

	resp.Document.AddressProof = model.DocumentLink{
		Url:         addViewUrl,
		DownloadUrl: addDownloadUrl,
		FileSize:    addressDocSize,
	}

	gstViewUrl, _ := service.GetFileURL(brandDetails.GstDocumant, 30)
	gstDownloadUrl, _ := service.CreateDownloadURL(brandDetails.GstDocumant, brandDetails.GstDocumentName, 30)
	gstressDocSize, _ := service.GetFileInfo(brandDetails.GstDocumant)

	resp.Document.GstDocument = model.DocumentLink{
		Url:         gstViewUrl,
		DownloadUrl: gstDownloadUrl,
		FileSize:    gstressDocSize,
	}

	panViewUrl, _ := service.GetFileURL(brandDetails.PanDocument, 30)
	panDownloadUrl, _ := service.CreateDownloadURL(brandDetails.PanDocument, brandDetails.PanDocumantName, 30)
	panressDocSize, _ := service.GetFileInfo(brandDetails.PanDocument)

	resp.Document.PanDocument = model.DocumentLink{
		Url:         panViewUrl,
		DownloadUrl: panDownloadUrl,
		FileSize:    panressDocSize,
	}

	return resp
}

func BrandInformation(db *gorm.DB, reqVal model.BrandInformationReq) model.GetBrandInformationRes {
	log := logger.InitLogger()
	var brandDetails model.GetBrandInformationFromDb
	err := db.Raw(query.GetBrandDataFromDbQuery, reqVal.BrandId).
		Scan(&brandDetails).Error
	if err != nil {
		log.Error("Error in Getting the Brand Information: " + err.Error())
		return model.GetBrandInformationRes{
			Status:  false,
			Message: "Failed to Get the Brand Information Data",
		}
	}

	response := MapBrandDetails(brandDetails)

	return model.GetBrandInformationRes{
		Status:    true,
		Message:   "Brand registration data fetched successfully",
		BrandData: response,
	}

}

func generatePassword(length int) (string, error) {
	// Allowed characters (letters, numbers, and special chars)
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"!@#$%^&*()-_=+<>?/{}[]|"

	password := make([]byte, length)

	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]
	}

	return string(password), nil
}
func BrandStatusUpdate(db *gorm.DB, reqVal model.BrandStatusUpdateReq) model.BrandStatusUpdateRes {
	log := logger.InitLogger()
	var brandDetails model.BrandStatusUpdateResFromDb
	date := getDate.GetCurrentDate("")

	// ✅ Brand Approved
	if reqVal.BrandApprove {
		// Generate random password
		pass, err := generatePassword(6)
		if err != nil {
			log.Error("Error generating password: " + err.Error())
			return model.BrandStatusUpdateRes{
				Status:  false,
				Message: "Password generation failed",
			}
		}

		// Hash password
		hashPassword, hashPasswordErr := becrypt.HashPassword(pass)
		if hashPasswordErr != nil {
			log.Error("ERROR: Failed to hash password: " + hashPasswordErr.Error())
			return model.BrandStatusUpdateRes{
				Status:  false,
				Message: "Something went wrong, try again",
			}
		}

		// Update brand status in DB
		err = db.Raw(
			query.UpdateBrandApprovalStatusQuery,
			reqVal.BrandId,
			4,            // ✅ Status ID: Approved
			"",           // Status description
			pass,         // plain password
			hashPassword, // hashed password
			date,
		).Scan(&brandDetails).Error

		if err != nil {
			log.Error("Error in updating brand status: " + err.Error())
			return model.BrandStatusUpdateRes{
				Status:  false,
				Message: "Failed to update brand status",
			}
		}

		// Prepare mail content
		subject := fmt.Sprintf("Congratulations %s - Your Brand is Approved!", brandDetails.BrandName)
		webSiteUrl := os.Getenv("BRAND_SELLER_LOGIN_URL")
		loginURL := webSiteUrl

		htmlContent := fmt.Sprintf(`
			<html>
				<body style="font-family: Arial, sans-serif; line-height: 1.6;">
					<h2>Hello %s,</h2>
					<p>We’re excited to inform you that your brand <b>%s</b> has been <span style="color:green;"><b>approved</b></span> in the Nivas portal!</p>
					
					<p>You can now log in using the credentials below:</p>
					
					<table style="border-collapse: collapse; margin: 10px 0;">
						<tr>
							<td style="padding: 8px; border: 1px solid #ddd;"><b>Email</b></td>
							<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
						</tr>
						<tr>
							<td style="padding: 8px; border: 1px solid #ddd;"><b>Password</b></td>
							<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
						</tr>
					</table>
					
					<p>
						<a href="%s" 
						   style="display:inline-block; padding:10px 20px; background:#4CAF50; 
								  color:#fff; text-decoration:none; border-radius:5px;">
							Login to Nivas Portal
						</a>
					</p>
					
					<p>If the button doesn’t work, copy and paste this link into your browser:</p>
					<p><a href="%s">%s</a></p>
					
					<br>
					<p>Best Regards,<br>Brand Support Team</p>
				</body>
			</html>
		`, brandDetails.ContactPerson, brandDetails.BrandName, brandDetails.BrandEmail, pass, loginURL, loginURL, loginURL)

		// Send mail
		mailSent := mailService.MailService(brandDetails.BrandEmail, htmlContent, subject)
		if !mailSent {
			log.Error("Mail sending failed for brand: " + brandDetails.BrandName)
			return model.BrandStatusUpdateRes{
				Status:  false,
				Message: "Brand approved, but mail sending failed",
			}
		}

		return model.BrandStatusUpdateRes{
			Status:  true,
			Message: "Brand approved and mail sent successfully",
		}
	}

	// ❌ Brand Rejected
	if !reqVal.BrandApprove {
		err := db.Raw(
			query.UpdateBrandRejectedStatus,
			reqVal.BrandId,
			5, // ✅ Status ID: Rejected
			reqVal.Reason,
			date,
		).Scan(&brandDetails).Error

		if err != nil {
			log.Error("Error in updating brand rejection status: " + err.Error())
			return model.BrandStatusUpdateRes{
				Status:  false,
				Message: "Failed to update brand rejection status",
			}
		}

		subject := fmt.Sprintf("Update on Your Brand Application - %s", brandDetails.BrandName)

		htmlContent := fmt.Sprintf(`
			<html>
				<body style="font-family: Arial, sans-serif; line-height: 1.6;">
					<h2>Hello %s,</h2>
					<p>We regret to inform you that your brand <b>%s</b> has been <span style="color:red;"><b>rejected</b></span> in the Nivas portal.</p>
					
					<p><b>Reason for Rejection:</b></p>
					<p style="padding:10px; background:#f8d7da; border:1px solid #f5c2c7; border-radius:5px; color:#721c24;">
						%s
					</p>
					
					<p>If you believe this is a mistake or would like to reapply, please reach out to our support team for further assistance.</p>
					
					<br>
					<p>Best Regards,<br>Brand Support Team</p>
				</body>
			</html>
		`, brandDetails.ContactPerson, brandDetails.BrandName, reqVal.Reason)

		mailSent := mailService.MailService(brandDetails.BrandEmail, htmlContent, subject)
		if !mailSent {
			log.Error("Mail sending failed for brand: " + brandDetails.BrandName)
			return model.BrandStatusUpdateRes{
				Status:  false,
				Message: "Brand rejected, but mail sending failed",
			}
		}
	}

	return model.BrandStatusUpdateRes{
		Status:  true,
		Message: "Brand status updated successfully",
	}
}
