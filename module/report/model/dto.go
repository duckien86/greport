package reportmodel

import "time"

type MsgLogFilter struct {
	StartDate  string   `json:"start_date" validate:"required,datetime=2006-01-02" form:"start_date"`
	EndDate    string   `json:"end_date" validate:"required,datetime=2006-01-02" form:"end_date"`
	KeyWord    string   `json:"key_word" validate:"omitempty,alphanum" form:"key_word"`
	Channel    string   `json:"channel" validate:"omitempty,alphanum" form:"channel"`
	CampaignId string   `json:"campaignId" validate:"omitempty,alphanum" form:"campaign_id"`
	TemplateId string   `json:"template_id" validate:"omitempty,alphanum" form:"template_id"`
	Sender     string   `json:"sender" validate:"omitempty,alphanum" form:"sender"`
	Recipient  string   `json:"recipient" validate:"omitempty,alphanum" form:"recipient"`
	MsgId      string   `json:"msg_id" validate:"omitempty,alphanum" form:"msg_id"`
	Telcos     []string `json:"telcos" validate:"omitempty,dive,alphanum" form:"telcos"`
	PoIds      []string `json:"poIds" validate:"omitempty,dive,alphanum" form:"po_ids"`
	AppNames   []string `json:"appNames" validate:"omitempty,dive,alphanum" form:"app_names"`
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
