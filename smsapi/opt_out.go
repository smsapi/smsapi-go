package smsapi

import (
	"context"
	"fmt"
)

const optOutsApiPath = "/opt_outs"

type OptOutApi struct {
	client *Client
}

type OptOut struct {
	Id          string `json:"id"`
	PhoneNumber int64  `json:"phoneNumber"`
	Date        string `json:"date"`
}

type OptOutCollection struct {
	Size       int       `json:"size"`
	Collection []*OptOut `json:"collection"`
}

type OptOutCollectionFilters struct {
	PaginationFilters
	PhoneNumber string `url:"phone_number,omitempty"`
}

type OptOutSettings struct {
	Brand string `json:"brand,omitempty"`
}

func (api *OptOutApi) List(ctx context.Context, filters *OptOutCollectionFilters) (*OptOutCollection, error) {
	result := new(OptOutCollection)
	uri, _ := addQueryParams(optOutsApiPath, filters)
	err := api.client.Get(ctx, uri, result)
	return result, err
}

func (api *OptOutApi) Delete(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s", optOutsApiPath, id)
	return api.client.Delete(ctx, uri)
}

func (api *OptOutApi) GetSettings(ctx context.Context) (*OptOutSettings, error) {
	result := new(OptOutSettings)
	err := api.client.Get(ctx, "/opt_outs/settings", result)
	return result, err
}

func (api *OptOutApi) UpdateSettings(ctx context.Context, settings *OptOutSettings) (*OptOutSettings, error) {
	result := new(OptOutSettings)
	err := api.client.Put(ctx, "/opt_outs/settings", result, settings)
	return result, err
}
