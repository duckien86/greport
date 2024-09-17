package reportmodel

import "time"

type MsgLogFilter struct {
	StartDate  string   `json:"start_date" validate:"required,datetime=2006-01-02"  `
	EndDate    string   `json:"end_date" validate:"required,datetime=2006-01-02" `
	KeyWord    string   `json:"key_word,omitempty"`
	Channel    string   `json:"channel,omitempty"`
	CampaignId string   `json:"campaignId,omitempty"`
	TemplateId string   `json:"template_id,omitempty"`
	Sender     string   `json:"sender,omitempty"`
	Recipient  string   `json:"recipient,omitempty"`
	MsgId      string   `json:"msg_id,omitempty"`
	Telcos     []string `json:"telcos,omitempty"`
	PoIds      []string `json:"poIds,omitempty"`
	AppNames   []string `json:"appNames,omitempty"`
}

type MsgLogResponse struct {
	MessageId      string     `json:"message_id"`
	ContactId      string     `json:"contact_id"`
	CampaignId     uint32     `json:"campaign_id"`
	CampaignName   string     `json:"campaign_name"`
	TemplateId     string     `json:"template_id"`
	Channel        string     `json:"channel"`
	App            string     `json:"app"`
	Sender         string     `json:"sender"`
	Recipient      string     `json:"recipient"`
	Telco          string     `json:"telco"`
	TimeSent       *time.Time `json:"time_sent"`
	SentStatus     string     `json:"sent_status"`
	DeliveryStatus string     `json:"delivery_status"`
	TimeDelivery   *time.Time `json:"time_delivery"`
	OpenStatus     string     `json:"open_status"`
	ErrorCode      string     `json:"error_code"`
	ErrorMsg       string     `json:"error_msg"`
	MTCount        string     `json:"mt_count"`
	PO             string     `json:"po"`
	IsFallback     string     `json:"is_fallback"`
	MsgBody        string     `json:"msg_body"`
	MsgType        string     `json:"msg_type"`
	MsgOption      string     `json:"msg_option"`
	SentBy         string     `json:"sent_by"`
}
