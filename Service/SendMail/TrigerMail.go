package TrigerMailService

import (
	"fmt"
	logger "nivasBrandRegistrationBackend/Helper/Logger"
	mailService "nivasBrandRegistrationBackend/Helper/MailService"
	TrigerMailModel "nivasBrandRegistrationBackend/Model/SendMail"
	"os"
)

func SendMail(reqVal TrigerMailModel.SendMailRequest) TrigerMailModel.SendMailReponce {
	log := logger.InitLogger()
	mailId := os.Getenv("BRAND_MAILID")
	subject := fmt.Sprintf(reqVal.Subject)
	htmlContent := fmt.Sprintf(reqVal.Content)
	mailSent := mailService.MailService(mailId, htmlContent, subject)
	if !mailSent {
		log.Error("Failed to send the mail ")
		return TrigerMailModel.SendMailReponce{
			Status:  false,
			Message: "but mail sending failed",
		}
	}
	log.Info("Brand application created and mail sent successfully to: " + mailId)
	return TrigerMailModel.SendMailReponce{
		Status:  true,
		Message: "Mail Send successfully",
	}

}
