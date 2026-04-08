package smsapi

import (
	"context"
	"fmt"
)

const callbacksApiPath = "/callbacks"

type CallbacksApi struct {
	client *Client
}

// Callback represents a single callback registration.
// Fields ReceiverType and ReceiverNumber are only used by `sms_mo` / `mms_mo` types.
type Callback struct {
	Id             string `json:"id,omitempty"`
	Url            string `json:"url"`
	Type           string `json:"type"`
	Active         bool   `json:"active,omitempty"`
	Invalid        *bool  `json:"invalid,omitempty"`
	ReceiverType   string `json:"receiver_type,omitempty"`
	ReceiverNumber string `json:"receiver_number,omitempty"`
}

type CallbackCollection struct {
	Size       int         `json:"size"`
	Collection []*Callback `json:"collection"`
}

type UpdateCallback struct {
	Url string `json:"url"`
}

type CallbackTestResult struct {
	TestResult         bool   `json:"test_result"`
	ConnectionFailed   bool   `json:"connection_failed"`
	InvalidEncoding    bool   `json:"invalid_encoding"`
	HttpResponseStatus *int   `json:"http_response_status,omitempty"`
	HttpResponseBody   string `json:"http_response_body,omitempty"`
}

func (api *CallbacksApi) List(ctx context.Context) (*CallbackCollection, error) {
	result := new(CallbackCollection)
	err := api.client.Get(ctx, callbacksApiPath, result)
	return result, err
}

func (api *CallbacksApi) Get(ctx context.Context, id string) (*Callback, error) {
	result := new(Callback)
	uri := fmt.Sprintf("%s/%s", callbacksApiPath, id)
	err := api.client.Get(ctx, uri, result)
	return result, err
}

func (api *CallbacksApi) Create(ctx context.Context, callback *Callback) (*Callback, error) {
	result := new(Callback)
	err := api.client.Post(ctx, callbacksApiPath, result, callback)
	return result, err
}

func (api *CallbacksApi) Update(ctx context.Context, id, url string) (*Callback, error) {
	result := new(Callback)
	uri := fmt.Sprintf("%s/%s", callbacksApiPath, id)
	err := api.client.Put(ctx, uri, result, &UpdateCallback{Url: url})
	return result, err
}

func (api *CallbacksApi) Delete(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s", callbacksApiPath, id)
	return api.client.Delete(ctx, uri)
}

func (api *CallbacksApi) Activate(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s/commands/activate", callbacksApiPath, id)
	return api.client.Put(ctx, uri, nil, nil)
}

func (api *CallbacksApi) Deactivate(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s/commands/deactivate", callbacksApiPath, id)
	return api.client.Put(ctx, uri, nil, nil)
}

func (api *CallbacksApi) Test(ctx context.Context, id string) (*CallbackTestResult, error) {
	result := new(CallbackTestResult)
	uri := fmt.Sprintf("%s/%s/commands/test", callbacksApiPath, id)
	err := api.client.Get(ctx, uri, result)
	return result, err
}
