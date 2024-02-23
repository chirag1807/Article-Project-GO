package utils

import (
	"articleproject/config"

	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/rest/api/v2010"
)

//this is only for learning purpose, this is not related to this article project.
// for detailed information go to https://github.com/twilio/twilio-go.

func MakeVoiceVall(to_mobile_no, voice_url string) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.Config.Twillio.AccountSID,
		Password: config.Config.Twillio.AuthToken,
	})

	from_mobile_no := config.Config.Twillio.FromMobileNo

	params_call := &openapi.CreateCallParams{
		To: &to_mobile_no,
		From: &from_mobile_no,
		Url: &voice_url, //in this url, you can pass any mp3 recording (url only) which will going to play when user receive this call.
	}
	_, err := client.Api.CreateCall(params_call)
	
	if err != nil {
		return
	}
}
