package brandRegistrationService

import (
	"bytes"
	"fmt"
	accesstoken "nivasBrandRegistrationBackend/Helper/AccessToken"
	getDate "nivasBrandRegistrationBackend/Helper/DateFormat"
	logger "nivasBrandRegistrationBackend/Helper/Logger"
	mailService "nivasBrandRegistrationBackend/Helper/MailService"
	BrandInviteMailTemplate "nivasBrandRegistrationBackend/Helper/MailTemplate"
	brandRegistrationModel "nivasBrandRegistrationBackend/Model/BrandOnBoading"
	brandRegistrationQuery "nivasBrandRegistrationBackend/Query/BrandOnBoading"
	"os"
	"text/template"

	"gorm.io/gorm"
)

func GenerateBrandRegisterUrl(db *gorm.DB, reqVal brandRegistrationModel.BrandRegistrationUrlRequest) brandRegistrationModel.BrandEmailCheckRes {
	log := logger.InitLogger()
	webSiteUrl := os.Getenv("BRAND_REGISTRATION_URL")

	var brandDetails []brandRegistrationModel.BrandMailCheckQueryResponse
	err := db.Raw(brandRegistrationQuery.CheckBrandEmail, reqVal.MailId).
		Scan(&brandDetails).Error
	if err != nil {
		log.Error("Brand Email Unique Check Query Failed: " + err.Error())
		return brandRegistrationModel.BrandEmailCheckRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}

	var applicationId brandRegistrationModel.BrandApplicationIdResponse
	var brandEmail, brandName, brandOwnerName, Message string

	if len(brandDetails) != 0 {
		fmt.Println("Brand Details:", brandDetails)

		applicationId.ApplicationCustId = brandDetails[0].ApplicationCustId
		brandName = brandDetails[0].BrandName
		brandOwnerName = brandDetails[0].BrandName // (maybe should be BrandOwnerName if you have that field later?)
		brandEmail = brandDetails[0].MailId
		Message = "Brand Already Registered, Mail Resend To that Brand"
	}

	// 2. Insert new brand application
	if len(brandDetails) == 0 {
		var date = getDate.GetCurrentDate("")
		fmt.Println(date)
		err := db.Raw(brandRegistrationQuery.CreateNewApplication, reqVal.BrandName, reqVal.MailId, reqVal.CustomerName, date).
			Scan(&applicationId).Error
		if err != nil {
			log.Error("Create New Application Insert Failed: " + err.Error())
			return brandRegistrationModel.BrandEmailCheckRes{
				Status:  false,
				Message: "Failed to create brand application",
			}
		}
		brandEmail = reqVal.MailId
		brandName = reqVal.BrandName
		brandOwnerName = reqVal.CustomerName
		Message = "Brand application created"

	}

	var token = accesstoken.CreateDayExpiringToken(30, applicationId.ApplicationCustId)

	registrationURL := webSiteUrl + "/#token:" + token
	subject := "Start your Journey With Us"
	brandMobileNumber := "+91 9797 787 989"
	brandMailId := os.Getenv("BRAND_MAILID")

	data := struct {
		ContactPerson   string
		BrandName       string
		RegistrationURL string
		BrandMailId     string
		BrandMobile     string
	}{
		ContactPerson:   brandOwnerName,
		BrandName:       brandName,
		RegistrationURL: registrationURL,
		BrandMailId:     brandMailId,
		BrandMobile:     brandMobileNumber,
	}

	tmpl, err := template.New("inviteMail").Parse(BrandInviteMailTemplate.InviteTemplate)
	if err != nil {
		log.Error("Error parsing email template: " + err.Error())
		return brandRegistrationModel.BrandEmailCheckRes{
			Status:  false,
			Message: "Template parsing error",
		}
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Error("Error executing email template: " + err.Error())
		return brandRegistrationModel.BrandEmailCheckRes{
			Status:  false,
			Message: "Template execution error",
		}
	}
	htmlContent := buf.String()

	// 4. Send mail
	mailSent := mailService.MailService(reqVal.MailId, htmlContent, subject)
	if !mailSent {
		log.Error("Mail sending failed for brand: " + brandName)
		return brandRegistrationModel.BrandEmailCheckRes{
			Status:  false,
			Message: Message + "but mail sending failed",
		}
	}

	log.Info("Brand application created and mail sent successfully to: " + brandEmail)
	return brandRegistrationModel.BrandEmailCheckRes{
		Status:  true,
		Message: Message + "successfully",
	}
}
