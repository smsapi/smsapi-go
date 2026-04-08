package smsapi

import (
	"context"
	"fmt"
	"net/http"
)

const smsTemplatesApiPath = "/sms/templates"

type SmsTemplatesApi struct {
	client *Client
}

type SmsTemplate struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name" url:"name"`
	Template  string `json:"template" url:"template"`
	Normalize bool   `json:"normalize,omitempty" url:"normalize,omitempty"`
}

type SmsTemplateCollection struct {
	Size       int            `json:"size"`
	Collection []*SmsTemplate `json:"collection"`
}

type AvailableSmsTemplate struct {
	Name      string `json:"name"`
	Template  string `json:"template"`
	Normalize bool   `json:"normalize,omitempty"`
}

type AvailableSmsTemplateCollection struct {
	Size       int                     `json:"size"`
	Collection []*AvailableSmsTemplate `json:"collection"`
}

func (api *SmsTemplatesApi) List(ctx context.Context) (*SmsTemplateCollection, error) {
	result := new(SmsTemplateCollection)
	err := api.client.Get(ctx, smsTemplatesApiPath, result)
	return result, err
}

func (api *SmsTemplatesApi) Get(ctx context.Context, id string) (*SmsTemplate, error) {
	result := new(SmsTemplate)
	uri := fmt.Sprintf("%s/%s", smsTemplatesApiPath, id)
	err := api.client.Get(ctx, uri, result)
	return result, err
}

func (api *SmsTemplatesApi) Create(ctx context.Context, template *SmsTemplate) (*SmsTemplate, error) {
	result := new(SmsTemplate)
	err := api.client.Urlencoded(ctx, http.MethodPost, smsTemplatesApiPath, result, template)
	return result, err
}

func (api *SmsTemplatesApi) Update(ctx context.Context, id string, template *SmsTemplate) (*SmsTemplate, error) {
	result := new(SmsTemplate)
	uri := fmt.Sprintf("%s/%s", smsTemplatesApiPath, id)
	err := api.client.Urlencoded(ctx, http.MethodPut, uri, result, template)
	return result, err
}

func (api *SmsTemplatesApi) Delete(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s", smsTemplatesApiPath, id)
	return api.client.Delete(ctx, uri)
}

// ListAvailable returns own templates and (optionally) ones shared with the main account.
func (api *SmsTemplatesApi) ListAvailable(ctx context.Context) (*AvailableSmsTemplateCollection, error) {
	result := new(AvailableSmsTemplateCollection)
	err := api.client.Get(ctx, "/sms/templates/available", result)
	return result, err
}
