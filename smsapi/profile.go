package smsapi

import (
	"context"
)

const (
	profileApiPath = "/profile"
)

type ProfileApi struct {
	client *Client
}

type ProfileDetailsResponse struct {
	Name        string  `json:"name,omitempty"`
	Email       string  `json:"email,omitempty"`
	Username    string  `json:"username,omitempty"`
	PhoneNumber string  `json:"phone_number,omitempty"`
	PaymentType string  `json:"payment_type,omitempty"`
	UserType    string  `json:"user_type,omitempty"`
	Points      float32 `json:"points,omitempty"`
}

func (accountApi *ProfileApi) Details(ctx context.Context) (*ProfileDetailsResponse, error) {
	var result = new(ProfileDetailsResponse)

	err := accountApi.client.Get(ctx, profileApiPath, result)

	return result, err
}

type Money struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type ProfilePrice struct {
	Price     *Money `json:"price,omitempty"`
	Country   string `json:"country,omitempty"`
	Network   string `json:"network,omitempty"`
	ChangedAt string `json:"changed_at,omitempty"`
}

type ProfilePricesResponse struct {
	Size       int             `json:"size"`
	Collection []*ProfilePrice `json:"collection"`
}

type ProfilePricesFilters struct {
	Type string `url:"type,omitempty"`
}

// Prices returns SMS/VMS/etc. pricing for this profile. `pricingType` selects the
// service (e.g. "pro", "eco", "sms", "2way", "vms", "hlr", "mms"); pass empty string for all.
func (accountApi *ProfileApi) Prices(ctx context.Context, pricingType string) (*ProfilePricesResponse, error) {
	var result = new(ProfilePricesResponse)

	uri, _ := addQueryParams("/profile/prices", &ProfilePricesFilters{Type: pricingType})

	err := accountApi.client.Get(ctx, uri, result)

	return result, err
}
