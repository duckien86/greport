package sender

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/rabbitmq/workqueues"
	"encoding/json"
	"errors"
)

const (
	SendTypeSms = "sms"
)

type SmsService struct {
	config map[string]string
}
type SmsData struct {
	To      string
	Content string
}

func NewSms() *SmsService {
	config := map[string]string{}
	return &SmsService{config: config}
}

func (s *SmsService) Send(msg SmsData, isAsync bool) error {
	if isAsync {
		msgBody, _ := json.Marshal(msg)
		workqueues.Publish(SendTypeSms, string(msgBody))
	}
	return nil
}

var (
	ErrSendSmsFail = common.NewCustomError(
		errors.New("send sms fail"),
		"send sms fail",
		"ErrSendSmsFail",
	)
)
