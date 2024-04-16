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
