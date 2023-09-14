package smsapi

import "context"

type MmsApi struct {
	client *Client
}

type Mms struct {
	To           string     `json:"to,omitempty"`
	Group        string     `json:"group,omitempty"`
	Subject      string     `json:"subject,omitempty"`
	Message      *SMIL      `json:"smil,omitempty"`
	Date         *Timestamp `json:"date,omitempty"`
	DateValidate string     `json:"date_validate,omitempty"`
	Idx          string     `json:"idx,omitempty"`
	CheckIdx     bool       `json:"check_idx,omitempty"`
	NotifyUrl    string     `json:"notify_url,omitempty"`
	Test         bool       `json:"test,omitempty"`
}

type MmsResponse struct {
	Id              string     `json:"id,omitempty"`
	Points          string     `json:"points,omitempty"`
	Number          string     `json:"number,omitempty"`
	DateSent        *Timestamp `json:"date_sent,omitempty"`
	SubmittedNumber string     `json:"submitted_number,omitempty"`
	Status          string     `json:"status,omitempty"`
	Idx             string     `json:"idx,omitempty"`
	Error           string     `json:"error,omitempty"`
}

type MmsCollectionResponse struct {
	Count int `json:"count"`

	Collection []*MmsResponse `json:"list"`
}

func (mmsApi *MmsApi) SendRaw(ctx context.Context, mms *Mms) (*MmsCollectionResponse, error) {
	var result = new(MmsCollectionResponse)

	err := mmsApi.client.LegacyPost(ctx, "/mms.do", result, mms)

	return result, err
}

func (mmsApi *MmsApi) Send(ctx context.Context, to, subject, image string) (*MmsCollectionResponse, error) {
	smil := NewSMIL()
	smil.AddImage(image)

	mms := &Mms{
		To:      to,
		Subject: subject,
		Message: smil,
	}

	return mmsApi.SendRaw(ctx, mms)
}

func (mmsApi *MmsApi) Schedule(ctx context.Context, to, subject, image string, sendAt *Timestamp) (*MmsCollectionResponse, error) {
	smil := NewSMIL()
	smil.AddImage(image)

	mms := &Mms{
		To:      to,
		Message: smil,
		Subject: subject,
		Date:    sendAt,
	}

	return mmsApi.SendRaw(ctx, mms)
}

func (mmsApi *MmsApi) SendToGroup(ctx context.Context, group, subject, image string) (*MmsCollectionResponse, error) {
	smil := NewSMIL()
	smil.AddImage(image)

	mms := &Mms{
		Group:   group,
		Subject: subject,
		Message: smil,
	}

	return mmsApi.SendRaw(ctx, mms)
}

type MmsRemoveResponse struct {
	Count int `json:"count"`

	Collection []*struct {
		Id string `json:"id,omitempty"`
	} `json:"list"`
}

func (mmsApi *MmsApi) RemoveScheduled(ctx context.Context, id string) (*MmsRemoveResponse, error) {
	var result = new(MmsRemoveResponse)

	payload := struct {
		SchDel string `json:"sch_del"`
	}{SchDel: id}

	err := mmsApi.client.LegacyPost(ctx, "/mms.do", result, payload)

	return result, err
}

func (mmsApi *MmsApi) Get(ctx context.Context, id string) (*MmsResponse, error) {
	var result = new(MmsResponse)

	v := struct {
		Status string `url:"status"`
	}{Status: id}

	uri, _ := addQueryParams("/mms.do", v)

	err := mmsApi.client.LegacyGet(ctx, uri, result)

	return result, err
}
