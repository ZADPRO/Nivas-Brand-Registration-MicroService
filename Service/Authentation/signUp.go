package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	signUpModelModel "nivasBrandRegistrationBackend/Model/Authentation"
	"os"
	"sort"

	"github.com/joho/godotenv"
	"github.com/nyaruka/phonenumbers"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"google.golang.org/api/idtoken"
)

func SignUpService(reqVal signUpModelModel.GoogleToken) signUpModelModel.SignUpResponse {
	fmt.Println(reqVal)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientId := os.Getenv("GOOGLE_CLIENT_ID")

	fmt.Println(clientId)
	payload, err := idtoken.Validate(
		context.Background(),
		reqVal.GoogleToken, clientId)

	fmt.Println(payload)
	fmt.Println(err)

	if err != nil {
		return signUpModelModel.SignUpResponse{
			Status:  false,
			Message: "Error in Google login",
		}
	}

	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	picture, _ := payload.Claims["picture"].(string)

	// if payload.

	return signUpModelModel.SignUpResponse{
		Status:  true,
		Message: "Google Login Successfully",
		Name:    name,
		Mail:    email,
		Profile: picture,
	}
}

func CountryCode() signUpModelModel.CountryInfoResponse {
	var result []signUpModelModel.CountryInfo
	regionNamer := display.English.Regions()

	for region := range phonenumbers.GetSupportedRegions() {
		// get dial code
		code := phonenumbers.GetCountryCodeForRegion(region)

		// try example number
		example := phonenumbers.GetExampleNumber(region)
		var minLen, maxLen int
		if example != nil {
			nsn := fmt.Sprintf("%d", example.GetNationalNumber())
			minLen = len(nsn)
			maxLen = len(nsn)
		}

		// get full country name
		langRegion := language.MustParseRegion(region)
		countryName := regionNamer.Name(langRegion)

		result = append(result, signUpModelModel.CountryInfo{
			CountryCode: region,      // e.g. "IN"
			CountryName: countryName, // e.g. "India"
			DialCode:    fmt.Sprintf("+%d", code),
			NSNMin:      minLen,
			NSNMax:      maxLen,
		})
	}

	// ✅ Sort by CountryName (ascending)
	sort.Slice(result, func(i, j int) bool {
		return result[i].CountryName < result[j].CountryName
	})

	return signUpModelModel.CountryInfoResponse{
		Status:      true,
		Message:     "Country Codes Passed Successfully",
		CountryData: result,
	}
}

func MobileNumberValidationService(reqVal signUpModelModel.MobileNumberValidationRequest) signUpModelModel.MobileNumberValidationResponse {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	twilioNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(reqVal.MobileNumber) // recipient number with country code
	params.SetFrom(twilioNumber)      // your Twilio number

	n, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		panic(err)
	}
	otp := 100000 + n.Int64() // ensures 6-digit
	fmt.Println("Random 6-digit number:", otp)

	params.SetBody(fmt.Sprintf("Use this OTP %d to login in Nivas", otp))
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Printf("Error sending message: %v", err)
		return signUpModelModel.MobileNumberValidationResponse{
			Status:  false,
			Message: "Error in Sending the OTP",
			Code:    otp,
		}
	}

	fmt.Printf("Message sent successfully! SID: %s\n", *resp.Sid)

	return signUpModelModel.MobileNumberValidationResponse{
		Status:  true,
		Message: "the Sms Send Successfully",
		Code:    otp,
	}
}
