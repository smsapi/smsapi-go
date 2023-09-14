package smsapi

import "context"

type Sms struct {
	To      string `json:"to,omitempty"`
	Message string `json:"message,omitempty"`
	From    string `json:"from,omitempty"`

	Group           string     `json:"group,omitempty"`
	Flash           bool       `json:"flash,omitempty"`
	Fast            bool       `json:"fast,omitempty"`
	Test            bool       `json:"test,omitempty"`
	Encoding        string     `json:"encoding,omitempty"`
	Details         bool       `json:"details,omitempty"`
	Date            *Timestamp `json:"date,omitempty"`
	Datacoding      string     `json:"datacoding,omitempty"`
	Udh             string     `json:"udh,omitempty"`
	SkipForeign     bool       `json:"skip_foreign,omitempty"`
	AllowDuplicates bool       `json:"allow_duplicates,omitempty"`
	Idx             int        `json:"idx,omitempty"`
	CheckIdx        bool       `json:"check_idx,omitempty"`
	Nounicode       bool       `json:"nounicode,omitempty"`
	Normalize       bool       `json:"normalize,omitempty"`
	PartnerId       bool       `json:"partner_id,omitempty"`
	MaxParts        int        `json:"max_parts,omitempty"`
	ExpirationDate  int        `json:"expiration_date,omitempty"`
	DiscountGroup   string     `json:"discount_group,omitempty"`
	NotifyUrl       string     `json:"notify_url,omitempty"`

	Template      string `json:"template,omitempty"`
	MessageParam1 string `json:"param1,omitempty"`
	MessageParam2 string `json:"param2,omitempty"`
	MessageParam3 string `json:"param3,omitempty"`
	MessageParam4 string `json:"param4,omitempty"`
}

type SmsResultCollection struct {
	Count int `json:"count"`

	Collection []*SmsResponse `json:"list"`
}

type SmsResponse struct {
	Id              string     `json:"id,omitempty"`
	Points          float32    `json:"points,omitempty"`
	Number          string     `json:"number,omitempty"`
	DateSent        *Timestamp `json:"date_sent,omitempty"`
	SubmittedNumber string     `json:"submitted_number,omitempty"`
	Status          string     `json:"status,omitempty"`
	Idx             string     `json:"idx,omitempty"`
	Error           string     `json:"error,omitempty"`
	Message         string     `json:"message,omitempty"`
	Length          int        `json:"length,omitempty"`
	Parts           int        `json:"parts,omitempty"`
}

type SmsApi struct {
	client *Client
}

func (smsApi *SmsApi) SendRaw(ctx context.Context, sms *Sms) (*SmsResultCollection, error) {
	var result = new(SmsResultCollection)

	err := smsApi.client.LegacyPost(ctx, "/sms.do", result, sms)

	return result, err
}

func (smsApi *SmsApi) Schedule(ctx context.Context, to, message string, from string, sendAt *Timestamp) (*SmsResultCollection, error) {
	sms := &Sms{
		To:      to,
		Message: message,
		From:    from,
		Date:    sendAt,
	}

	return smsApi.SendRaw(ctx, sms)
}

func (smsApi *SmsApi) Send(ctx context.Context, to, message string, from string) (*SmsResultCollection, error) {
	sms := &Sms{
		To:      to,
		Message: message,
		From:    from,
	}

	return smsApi.SendRaw(ctx, sms)
}

func (smsApi *SmsApi) SendFlash(ctx context.Context, to, message string, from string) (*SmsResultCollection, error) {
	sms := &Sms{
		To:      to,
		Message: message,
		From:    from,
		Flash:   true,
	}

	return smsApi.SendRaw(ctx, sms)
}

func (smsApi *SmsApi) SendToGroup(ctx context.Context, group, message string, from string) (*SmsResultCollection, error) {
	sms := &Sms{
		Group:   group,
		Message: message,
		From:    from,
	}

	return smsApi.SendRaw(ctx, sms)
}

type SmsRemoveResult struct {
	Count int `json:"count"`

	Collection []*struct {
		Id string `json:"id,omitempty"`
	} `json:"list"`
}

func (smsApi *SmsApi) RemoveScheduled(ctx context.Context, id string) (*SmsRemoveResult, error) {
	var result = new(SmsRemoveResult)

	payload := struct {
		SchDel string `json:"sch_del"`
	}{SchDel: id}

	err := smsApi.client.LegacyPost(ctx, "/sms.do", result, payload)

	return result, err
}

func (smsApi *SmsApi) Get(ctx context.Context, id string) (*SmsResponse, error) {
	var result = new(SmsResponse)

	v := struct {
		Status string `url:"status"`
	}{Status: id}

	uri, _ := addQueryParams("/sms.do", v)

	err := smsApi.client.LegacyGet(ctx, uri, result)

	return result, err
}
