package telegramsender

import (
	"log"

	"github.com/corporateanon/barker-worker/pkg/sender"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/go-resty/resty/v2"
)

type SenderImplTelegram struct {
	resty *resty.Client
}

func NewSenderImplTelegram(resty *resty.Client) sender.Sender {
	return &SenderImplTelegram{
		resty: resty,
	}
}

func (sender *SenderImplTelegram) Send(bot *types.Bot, campaign *types.Campaign, user *types.User) error {
	log.Printf("Sending campaign %s to user %v\n", campaign.Title, user)

	payload := createSendMessagePayload(user, campaign)

	res, err := sender.resty.R().
		SetBody(payload).
		SetPathParams(map[string]string{
			"Token": bot.Token,
		}).
		SetError(&ErrorResponse{}).
		Post("https://api.telegram.org/bot{Token}/sendMessage")

	if err != nil {
		return err
	}
	if httpErr := res.Error(); httpErr != nil {
		log.Printf("Error %v\n", httpErr)
		return httpErr.(*ErrorResponse)
	}
	log.Println("ok")

	return nil
}
