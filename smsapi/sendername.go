package smsapi

import (
	"context"
	"fmt"
)

type SenderApi struct {
	client *Client
}

type Sender struct {
	Name string `json:"sender"`
}

type SenderResponse struct {
	IsDefault bool   `json:"is_default"`
	Name      string `json:"sender"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type SenderCollectionResponse struct {
	Size       int               `json:"size"`
	Collection []*SenderResponse `json:"collection"`
}

func (senderApi *SenderApi) Get(ctx context.Context, name string) (*SenderResponse, error) {
	uri := fmt.Sprintf("/sms/sendernames/%s", name)

	var result = new(SenderResponse)

	err := senderApi.client.Get(ctx, uri, result)

	return result, err
}

func (senderApi *SenderApi) GetAll(ctx context.Context) (*SenderCollectionResponse, error) {
	var result = new(SenderCollectionResponse)

	err := senderApi.client.Get(ctx, "/sms/sendernames", result)

	return result, err
}

func (senderApi *SenderApi) Create(ctx context.Context, name string) (*SenderResponse, error) {
	sender := &Sender{
		Name: name,
	}

	var result = new(SenderResponse)

	err := senderApi.client.Post(ctx, "/sms/sendernames", result, sender)

	return result, err
}

func (senderApi *SenderApi) Delete(ctx context.Context, name string) error {
	uri := fmt.Sprintf("/sms/sendernames/%s", name)

	err := senderApi.client.Delete(ctx, uri)

	return err
}

func (senderApi *SenderApi) Activate(ctx context.Context, name, code string) error {
	uri := fmt.Sprintf("/sms/sendernames/%s/commands/activate", name)

	payload := struct {
		Code string `json:"code"`
	}{
		Code: code,
	}

	err := senderApi.client.Put(ctx, uri, nil, &payload)

	return err
}

func (senderApi *SenderApi) MakeDefault(ctx context.Context, name string) error {
	uri := fmt.Sprintf("/sms/sendernames/%s/commands/make_default", name)

	err := senderApi.client.Post(ctx, uri, nil, nil)

	return err
}
