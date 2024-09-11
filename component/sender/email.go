package sender

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/rabbitmq/workqueues"
	"encoding/json"
	"errors"
)

const (
	SendTypeEmail = "email"
)

type EmailData struct {
	To      string
	Content string
}

type EmailService struct {
	config map[string]string
}

func NewEmail() *EmailService {
	config := map[string]string{}
	return &EmailService{config: config}
}

// isAsync: true is mean send email in queue task
// and false is mean send email immediately
func (s *EmailService) Send(msg EmailData, isAsync bool) error {
	// send
	if isAsync {
		msgBody, _ := json.Marshal(msg)
		workqueues.Publish(SendTypeEmail, string(msgBody))
	}
	return nil
}

var (
	ErrSendEmailFail = common.NewCustomError(
		errors.New("send sms fail"),
		"send sms fail",
		"ErrSendSmsFail",
	)
)
