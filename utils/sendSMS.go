package utils

import (
	"articleproject/config"

	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/rest/api/v2010"
)

//this is only for learning purpose, this is not related to this article project.
// for detailed information go to https://github.com/twilio/twilio-go.

func SendSMS(to_mobile_no, body string) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.Config.Twillio.AccountSID,
		Password: config.Config.Twillio.AuthToken,
	})

	from_mobile_no := config.Config.Twillio.FromMobileNo

	params_sms := &openapi.CreateMessageParams{
		To:   &to_mobile_no,
		From: &from_mobile_no,
		Body: &body,
	}
	_, err := client.Api.CreateMessage(params_sms)
	if err != nil {
		return
	}
}
