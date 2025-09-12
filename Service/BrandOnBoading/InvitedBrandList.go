package brandRegistrationService

import (
	"bytes"
	"fmt"
	accesstoken "nivasBrandRegistrationBackend/Helper/AccessToken"
	logger "nivasBrandRegistrationBackend/Helper/Logger"
	mailService "nivasBrandRegistrationBackend/Helper/MailService"
	BrandInviteMailTemplate "nivasBrandRegistrationBackend/Helper/MailTemplate"
	brandRegistrationModel "nivasBrandRegistrationBackend/Model/BrandOnBoading"
	brandRegistrationQuery "nivasBrandRegistrationBackend/Query/BrandOnBoading"
	"os"
	"text/template"

	"gorm.io/gorm"
)

func InvitedBrandList(db *gorm.DB) brandRegistrationModel.InvitedBrandListRes {
	log := logger.InitLogger()

	var brandDetails []brandRegistrationModel.InvitedBrandList
	err := db.Raw(brandRegistrationQuery.InvitedBrandListQuery).
		Scan(&brandDetails).Error
	if err != nil {
		log.Error("Error in Getting the Brand Invited List: " + err.Error())
		return brandRegistrationModel.InvitedBrandListRes{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}

	return brandRegistrationModel.InvitedBrandListRes{
		Status:    true,
		Message:   "Invited Brand List Passed Successfully",
		BrandList: brandDetails,
	}
}

func ReSendBrandInvite(db *gorm.DB, brandData brandRegistrationModel.ReInviteBrandRequest) brandRegistrationModel.ReInviteBrandResponce {
	log := logger.InitLogger()

	// Fetch brand details
	var brandDetails []brandRegistrationModel.InvitedBrandList
	err := db.Raw(brandRegistrationQuery.ReInviteBrandDataQuery, brandData.BrandId).
		Scan(&brandDetails).Error
	if err != nil {
		log.Error("Error in Getting Re Invite Call Brand Data: " + err.Error())
		return brandRegistrationModel.ReInviteBrandResponce{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}
	if len(brandDetails) == 0 {
		log.Error("No brand details found for BrandId")
		return brandRegistrationModel.ReInviteBrandResponce{
			Status:  false,
			Message: "No brand details found",
		}
	}

	// Generate registration link
	webSiteUrl := os.Getenv("BRAND_REGISTRATION_URL")
	token := accesstoken.CreateDayExpiringToken(30, brandDetails[0].BrandCustId)
	registrationURL := webSiteUrl + "/#token:" + token

	// Mail subject and sender details
	subject := "Start your Journey With Us"
	brandMobileNumber := "+91 9797 787 989"
	brandMailId := os.Getenv("BRAND_MAILID")

	// Debug prints
	fmt.Println("ContactPerson:", brandDetails[0].ContactPerson)
	fmt.Println("BrandName:", brandDetails[0].BrandName)
	fmt.Println("RegistrationURL:", registrationURL)
	fmt.Println("BrandMailId:", brandMailId)
	fmt.Println("BrandMobile:", brandMobileNumber)

	// Prepare template data
	data := struct {
		ContactPerson   string
		BrandName       string
		RegistrationURL string
		BrandMailId     string
		BrandMobile     string
	}{
		ContactPerson:   brandDetails[0].ContactPerson,
		BrandName:       brandDetails[0].BrandName,
		RegistrationURL: registrationURL,
		BrandMailId:     brandMailId,
		BrandMobile:     brandMobileNumber,
	}

	// Parse and execute template
	tmpl, err := template.New("inviteMail").Parse(BrandInviteMailTemplate.InviteTemplate)
	if err != nil {
		log.Error("Error parsing email template: " + err.Error())
		return brandRegistrationModel.ReInviteBrandResponce{
			Status:  false,
			Message: "Template parsing error",
		}
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Error("Error executing email template: " + err.Error())
		return brandRegistrationModel.ReInviteBrandResponce{
			Status:  false,
			Message: "Template execution error",
		}
	}
	htmlContent := buf.String()

	// Send mail
	mailSent := mailService.MailService(brandDetails[0].ContactEmail, htmlContent, subject)
	if !mailSent {
		log.Error("Mail sending failed for brand: " + brandDetails[0].BrandName)
		return brandRegistrationModel.ReInviteBrandResponce{
			Status:  false,
			Message: "Error In Sending mail",
		}
	}

	log.Info("Brand application link resent successfully: " + brandDetails[0].ContactEmail)
	return brandRegistrationModel.ReInviteBrandResponce{
		Status:  true,
		Message: "Invite Mail Resend Successfully",
	}
}
