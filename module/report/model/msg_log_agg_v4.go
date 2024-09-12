package reportmodel

import "time"

type MsgLogReq struct {
	StartDate  *time.Time `json:"start_date" binding:"required" time_format:"2006-01-02" time_utc:"7"`
	EndDate    *time.Time `json:"end_date" binding:"required" time_format:"2006-01-02" time_utc:"7"`
	KeyWord    string     `json:"key_word,omitempty"`
	Channel    string     `json:"channel,omitempty"`
	CampaignId string     `json:"campaignId,omitempty"`
	TemplateId string     `json:"template_id,omitempty"`
	Sender     string     `json:"sender,omitempty"`
	Recipient  string     `json:"recipient,omitempty"`
	MsgId      string     `json:"msg_id,omitempty"`
	Telcos     []string   `json:"telcos,omitempty"`
	PoIds      []string   `json:"poIds,omitempty"`
	AppNames   []string   `json:"appNames,omitempty"`
}
