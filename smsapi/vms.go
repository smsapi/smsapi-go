package smsapi

import "context"

const vmsApiPath = "/vms.do"

type VmsApi struct {
	client *Client
}

type Vms struct {
	To           string     `json:"to,omitempty"`
	Group        string     `json:"group,omitempty"`
	From         string     `json:"from,omitempty"`
	Tts          string     `json:"tts,omitempty"`
	File         string     `json:"file,omitempty"`
	TtsLector    string     `json:"tts_lector,omitempty"`
	Date         *Timestamp `json:"date,omitempty"`
	DateValidate string     `json:"date_validate,omitempty"`
	Try          string     `json:"try,omitempty"`
	Interval     string     `json:"interval,omitempty"`
	SkipGms      string     `json:"skip_gsm,omitempty"`
	Idx          string     `json:"idx,omitempty"`
	CheckIdx     string     `json:"check_idx,omitempty"`
	NotifyUrl    string     `json:"notify_url,omitempty"`
	Test         string     `json:"test,omitempty"`
}

type VmsResponse struct {
	Id              string     `json:"id,omitempty"`
	Points          float32    `json:"points,omitempty"`
	Number          string     `json:"number,omitempty"`
	DateSent        *Timestamp `json:"date_sent,omitempty"`
	SubmittedNumber string     `json:"submitted_number,omitempty"`
	Status          string     `json:"status,omitempty"`
	Idx             string     `json:"idx,omitempty"`
	Error           string     `json:"error,omitempty"`
}

type VmsCollectionResponse struct {
	Count int `json:"count"`

	Collection []*VmsResponse `json:"list"`
}

func (vmsApi *VmsApi) SendRaw(ctx context.Context, vms *Vms) (*VmsCollectionResponse, error) {
	var result = new(VmsCollectionResponse)

	err := vmsApi.client.LegacyPost(ctx, vmsApiPath, result, vms)

	return result, err
}

func (vmsApi *VmsApi) Send(ctx context.Context, to, message, from string) (*VmsCollectionResponse, error) {
	vms := &Vms{
		To:   to,
		Tts:  message,
		From: from,
	}

	return vmsApi.SendRaw(ctx, vms)
}

func (vmsApi *VmsApi) Schedule(ctx context.Context, to, message, from string, sendAt *Timestamp) (*VmsCollectionResponse, error) {
	vms := &Vms{
		To:   to,
		Tts:  message,
		From: from,
		Date: sendAt,
	}

	return vmsApi.SendRaw(ctx, vms)
}

func (vmsApi *VmsApi) SendToGroup(ctx context.Context, group, message, from string) (*VmsCollectionResponse, error) {
	vms := &Vms{
		Group: group,
		Tts:   message,
		From:  from,
	}

	return vmsApi.SendRaw(ctx, vms)
}

type VmsRemoveResponse struct {
	Count int `json:"count"`

	Collection []*struct {
		Id string `json:"id,omitempty"`
	} `json:"list"`
}

func (vmsApi *VmsApi) RemoveScheduled(ctx context.Context, id string) (*VmsRemoveResponse, error) {
	var result = new(VmsRemoveResponse)

	v := struct {
		SchDel string `url:"sch_del"`
	}{SchDel: id}

	uri, _ := addQueryParams(vmsApiPath, v)

	err := vmsApi.client.LegacyGet(ctx, uri, result)

	return result, err
}

func (vmsApi *VmsApi) Get(ctx context.Context, id string) (*VmsResponse, error) {
	var result = new(VmsResponse)

	v := struct {
		Status string `url:"status"`
	}{Status: id}

	uri, _ := addQueryParams(vmsApiPath, v)

	err := vmsApi.client.LegacyGet(ctx, uri, result)

	return result, err
}
